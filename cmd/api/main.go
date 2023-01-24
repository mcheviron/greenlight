package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "Port", 4000, "API server port")
	flag.StringVar(
		&cfg.env,
		"Environment",
		"development",
		"Environment (development|staging|production)",
	)
	flag.StringVar(
		&cfg.db.dsn,
		"DB-DSN",
		os.Getenv("GREENLIGHT_DB_DSN"),
		"PostgresSQL DSN",
	)
	flag.IntVar(
		&cfg.db.maxOpenConns,
		"DB-Max-Open-Conns",
		25,
		"PostgreSQL max open connections",
	)
	flag.IntVar(
		&cfg.db.maxIdleConns,
		"DB-Max-Idle-Conns",
		25,
		"PostgreSQL max idle connections",
	)
	flag.StringVar(
		&cfg.db.maxIdleTime,
		"DB-Max-Idle-Time",
		"15m",
		"PostgreSQL max connection idle time",
	)
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Printf("database connection pool established")

	app := application{config: cfg, logger: logger}

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}

	logger.Printf("Starting %s server on port %d", cfg.env, cfg.port)
	logger.Fatal(srv.ListenAndServe())
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	duration, err := time.ParseDuration(cfg.db.maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
