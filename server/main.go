package main

import (
	"github.com/go-redis/redis/v9"
	"log"
	"net/http"
	"os"
	"server/db"
	"server/handlers"
)

func main() {
	connStr := "user=postgres password=123456 dbname=dvdrental sslmode=disable"
	addr := "localhost:8080"

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	timeLog := log.New(os.Stderr, "TIME\t", log.Ldate|log.Ltime|log.Lshortfile)

	pg, err := db.OpenDB(connStr)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pg.Close()

	cl := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer cl.Close()

	app := handlers.NewApplication(errorLog, timeLog, pg, cl, &addr)

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	log.Println("Запуск веб-сервера на", addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
