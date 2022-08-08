package main

func main() {
	/*
		addr := os.Getenv("SERVER_URL")
		dsn := os.Getenv("DATABASE_URL")

		infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		timeLog := log.New(os.Stderr, "TIME\t", log.Ldate|log.Ltime|log.Lshortfile)

		db, err := db.OpenDB()
		if err != nil {
			errorLog.Fatal(err)
		}
		defer db.Close()

		app := handlers.NewApplication()

		srv := &http.Server{
			Addr:     *addr,
			ErrorLog: errorLog,
			Handler:  app.Routes(),
		}

		app.InfoLog().Printf("Запуск веб-сервера на %s", *addr)
		err = srv.ListenAndServe()
		app.ErrorLog().Fatal(err)

	*/
}
