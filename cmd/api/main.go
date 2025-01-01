package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env string

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
	flag.Parse()

	// initialize new structured logger
	logger := slog.New(slog.NewTextHandler(os.Stdout,nil))

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

	err:= srv.ListenAndServe()
	if err!=nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}