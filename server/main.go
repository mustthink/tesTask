package main

import (
	"log"
	"net/http"
	"os"
	"server/db"
	"server/handlers"
)

func main() {
	connStr := "user=postgres password=123456 dbname=dvdrental sslmode=disable"
	addr := os.Getenv("SERVER_URL")

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	timeLog := log.New(os.Stderr, "TIME\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := db.OpenDB(connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := handlers.NewApplication(errorLog, timeLog, db, &addr)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	log.Println("Запуск веб-сервера на %s", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
