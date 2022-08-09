package handlers

import "net/http"

func (app *application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/film", app.showFilm)
	mux.HandleFunc("/film/list", app.showFilmsWithR)
	mux.HandleFunc("/film/create", app.createFilm)

	return mux
}
