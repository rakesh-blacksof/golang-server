package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port    int
	env     string
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.StringVar(&cfg.env, "env", "development", "HTTP Runtime Env")
	flag.IntVar(&cfg.port, "port", 1000, "HTTP Run Port")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter:RPS")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter:Burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Rate limiter:Enabled")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	logger.Printf("Starting the server on port : %d", app.config.port)

	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}

}
