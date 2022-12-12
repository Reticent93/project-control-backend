package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"project-control-backend/internal/repository"
	"project-control-backend/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	Domain   string
	DSN      string
	DB       repository.DatabaseRepo
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	var app application
	
	//read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=management sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	
	//connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	
	fmt.Println("Starting server on port", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
