package main

import (
	"context"
	"database/sql"
	"flag"
	"sync"

	"log/slog"

	"os"
	"time"

	"github.com/half-blood-prince-2710/greenlight-GO-RestApi/internal/data"
	"github.com/half-blood-prince-2710/greenlight-GO-RestApi/internal/mailer"
	_ "github.com/lib/pq"
)

// SIGINT Interrupt from keyboard Ctrl+C Yes
// SIGQUIT Quit from keyboard Ctrl+\ Yes
// SIGKILL Kill process (terminate immediately) - No
// SIGTERM Terminate process in orderly manner - Yes
const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
	mailer mailer.Mailer
	wg sync.WaitGroup
}


func main() {
	var cfg config
	// configuration flags
	flag.IntVar(&cfg.port, "port", 5000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN"), "PostgreSQL Dsn")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(&cfg.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "PostgreSQL max connection idle time")
	// Create command line flags to read the setting values into the config struct.
	// Notice that we use true as the default for the 'enabled' setting?
	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "sandbox.smtp.mailtrap.io", "SMTP host")
flag.IntVar(&cfg.smtp.port, "smtp-port", 25, "SMTP port")
flag.StringVar(&cfg.smtp.username, "smtp-username", os.Getenv(SMTP_USERNAME), "SMTP username")
flag.StringVar(&cfg.smtp.password, "smtp-password", os.Getenv(SMTP_PASSWORD), "SMTP password")
flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Greenlight <no-reply@greenlight.alexedwards.net>", "SMTP sender")
	flag.Parse()

	// initialize new structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// establish database
	db, err := openDB(cfg)
	if err != nil {
		logger.Error(err.Error())
	}
	defer db.Close()
	logger.Info("database connection esatblished")

	// declaring instance of application struct

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
		mailer: mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
