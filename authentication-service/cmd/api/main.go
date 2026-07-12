package main

import (
	"broker-service/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const WebPort = "8081"

type Application struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")
	db := connectDB()
	if db == nil {
		log.Println("Error connecting to database")
	}

	app := Application{
		DB:     db,
		Models: data.New(db),
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", WebPort),
		Handler: app.Routes(),
	}

	log.Printf("Starting authentication service on port %s\n", WebPort)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(fmt.Errorf("error starting http server: %v", err))
	}
}

func connectDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for range 10 {
		db, err := sql.Open("pgx", dsn)
		if err == nil && db.Ping() == nil {
			log.Printf("Connected to database")
			return db
		}
		log.Printf("Waiting for postgres")
		time.Sleep(2 * time.Second)
	}
	return nil
}
