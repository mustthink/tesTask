package handlers

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
	"net/http"
	"runtime/debug"
	"server/data"
)

type application struct {
	errorLog *log.Logger
	timeLog  *log.Logger
	data     *data.Service
	url      *string
}

func NewApplication(errorLog *log.Logger, timeLog *log.Logger, db *sql.DB, cl *redis.Client, url *string) *application {
	return &application{
		errorLog: errorLog,
		timeLog:  timeLog,
		data: &data.Service{
			DB:       db,
			Cache:    map[string][]byte{},
			RedStore: cl,
		},
		url: url,
	}
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
