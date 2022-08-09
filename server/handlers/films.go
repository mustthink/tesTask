package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/data"
	"server/db"
	"strconv"
	"time"
)

func (app *application) createFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	newFilm := &data.Film{
		Title:            "Shrek",
		Description:      "Somebody was told me",
		Release_year:     2001,
		Language_id:      1,
		Rental_duration:  1,
		Rental_rate:      1,
		Length:           120,
		Replacement_cost: 1,
		Rating:           "R",
		Last_update:      time.Now(),
		Special_features: "{Deleted scenes}",
		Fulltext:         "123",
	}

	err := app.data.InsertFilm(newFilm)
	if err != nil {
		app.serverError(w, err)
		return
	}

}

func (app *application) showFilm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	incache := true
	s := &data.Film{}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	title := ""

	if id < 1 {
		title = r.URL.Query().Get("title")
		if title == "" {
			app.notFound(w)
			return
		}
		s, err = app.data.CheckCache(title)
		if err != nil {
			app.serverError(w, err)
		}
		if s == nil {
			s, err = app.data.GetFilmViaTitle(title)
			incache = false
		}
	} else {
		if err != nil {
			app.serverError(w, err)
		}
		s, err = app.data.CheckCache(string(id))
		if err != nil {
			app.serverError(w, err)
		}
		if s == nil {
			s, err = app.data.GetFilmViaID(id)
			incache = false
		}
	}

	if err != nil {
		if errors.Is(err, db.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	json_data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		app.errorLog.Println(err)
	}
	if id < 1 && !incache {
		app.data.AddToCache(json_data, title)
	} else if !incache {
		app.data.AddToCache(json_data, string(id))
	}

	fmt.Fprintf(w, "%v", string(json_data))
}

func (app *application) showFilmsWithR(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	rate, err := strconv.Atoi(r.URL.Query().Get("rating"))
	if err != nil {
		app.notFound(w)
	}

	s, err := app.data.GetFilmsGreaterThen(rate)
	if err != nil {
		app.serverError(w, err)
		return
	}
	json_data, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		app.serverError(w, err)
	}

	fmt.Fprintf(w, "%v", string(json_data))
}
