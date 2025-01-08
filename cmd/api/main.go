package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env string
	db struct {
		dsn string
	}

}

type application struct {
	config config
	logger *slog.Logger
}


func main() {
	var cfg config
	// configuration flags
	flag.IntVar(&cfg.port, "port",5000,"API server port")
	flag.StringVar(&cfg.env,"env","development","Environment(development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn",os.Getenv("GREENLIGHT_DB_DSN"),"PostgreSQL Dsn")
	flag.Parse()

	
	// initialize new structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))


	// establish database 
	db , err := openDB(cfg)
	if err !=nil {
		logger.Error(err.Error())
	}
	defer db.Close()
	logger.Info("database connection esatblished")

	// declaring instance of application struct

	app:= &application{
		config : cfg,
		logger : logger,
	}

	

	// declare http server

	srv := &http.Server{
		Addr : fmt.Sprintf(":%d",cfg.port),
		Handler : app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 5*time.Second,
		WriteTimeout: 10*time.Second,
		ErrorLog: slog.NewLogLogger(logger.Handler(),slog.LevelError),
	}

	//start http server
	logger.Info("starting server","addr",srv.Addr,"env",cfg.env)

	err = srv.ListenAndServe()
	if err!=nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}

func openDB(cfg config) (*sql.DB,error) {
	db , err := sql.Open("postgres",cfg.db.dsn)
	if err!= nil {
		return nil, err
	}
	ctx , cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}

	return db , nil
}