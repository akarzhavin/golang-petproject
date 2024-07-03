package main

import (
	"blog-service/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "8080"
const defaultDSN = "host=localhost port=5432 user=postgresql password=password dbname=postgresql sslmode=disable timezone=UTC connect_timeout=5"

var counts int64

type Config struct {
	DB     *sql.DB
	Models models.Models
}

func main() {
	log.Println("Starting authentication service ...")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// Set up config
	app := Config{
		DB:     conn,
		Models: models.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn, isEnvFound := os.LookupEnv("DSN")
	if !isEnvFound {
		dsn = defaultDSN
	}

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Panicln("Postgres not yet ready ...")
			counts++
		} else {
			log.Println("Connect to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds ...")
		time.Sleep(2 * time.Second)
		continue
	}
}
