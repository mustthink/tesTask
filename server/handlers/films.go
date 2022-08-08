package handlers

import (
	"net/http"
	"server/data"
	"time"
)

func (app *application) insertFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	newFilm := &data.Film{
		Title:            "Shrek",
		Description:      "Somebody was told me",
		Realease_year:    2001,
		Language_id:      1,
		Rental_duration:  1,
		Rental_rate:      1,
		Length:           120,
		Replacement_cost: 1,
		Rating:           10,
		Last_update:      time.Now(),
		Special_features: "some cool cartoon",
		Fulltext:         "foo",
	}

	err := app.data.InsertFilm(newFilm)
	if err != nil {
		app.serverError(w, err)
		return
	}

}
