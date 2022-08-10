package data

import (
	"database/sql"
	"encoding/json"
	"errors"
	"server/db"
	"strconv"
	"time"
)

/*
CREATE TABLE public.film (
    film_id integer DEFAULT nextval('public.film_film_id_seq'::regclass) NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    release_year public.year,
    language_id smallint NOT NULL,
    rental_duration smallint DEFAULT 3 NOT NULL,
    rental_rate numeric(4,2) DEFAULT 4.99 NOT NULL,
    length smallint,
    replacement_cost numeric(5,2) DEFAULT 19.99 NOT NULL,
    rating public.mpaa_rating DEFAULT 'G'::public.mpaa_rating,
    last_update timestamp without time zone DEFAULT now() NOT NULL,
    special_features text[],
    fulltext tsvector NOT NULL
);*/

type Film struct {
	Film_id          int       `json:"film_id"`
	Title            string    `json:"title" `
	Description      string    `json:"description"`
	Release_year     int       `json:"realease_year"`
	Language_id      int       `json:"language_id"`
	Rental_duration  int       `json:"rental_duration"`
	Rental_rate      float64   `json:"rental_rate"`
	Length           int       `json:"length"`
	Replacement_cost float64   `json:"replacement_cost"`
	Rating           string    `json:"rating"`
	Last_update      time.Time `json:"last_update"`
	Special_features string    `json:"special_features"`
	Fulltext         string    `json:"fulltext"`
}

func (m *Service) CheckCache(key string) (film *Film, err error) {
	jdata, err := m.ShowFromCache(key)
	if err != nil {
		return nil, err
	}
	if jdata != nil {
		err = json.Unmarshal(jdata, &film)
		if err != nil {
			return nil, err
		}
		return film, nil
	}
	return nil, nil
}

func (m *Service) InsertFilm(film *Film) error {

	stmt := `insert into film (title, description, release_year, language_id, rental_duration, rental_rate, length, replacement_cost, rating, last_update, special_features, fulltext) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := m.DB.Exec(stmt, film.Title, film.Description, film.Release_year, film.Language_id, film.Rental_duration, film.Rental_rate, film.Length, film.Replacement_cost, film.Rating, film.Last_update, film.Special_features, film.Fulltext)
	if err != nil {
		return err
	}

	return nil
}

func (m *Service) GetFilmViaID(id int) (*Film, error) {
	start := time.Now()
	film := &Film{}
	stmt := `select * from film where film_id = $1`

	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNoRecord
		} else {
			return nil, err
		}
	}
	err = m.timeLog(strconv.Itoa(id), time.Since(start).Microseconds(), "db")
	if err != nil {
		return nil, err
	}
	return film, nil
}

func (m *Service) GetFilmViaTitle(title string) (*Film, error) {
	start := time.Now()
	film := &Film{}
	stmt := `select * from film where title = $1`

	row := m.DB.QueryRow(stmt, title)

	err := row.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNoRecord
		} else {
			return nil, err
		}
	}
	err = m.timeLog(title, time.Since(start).Microseconds(), "db")
	if err != nil {
		return nil, err
	}
	return film, nil
}

func (m *Service) GetFilmsGreaterThen(rate int) ([]*Film, error) {

	stmt := `select * from film where rental_rate > $1`

	rows, err := m.DB.Query(stmt, rate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var films []*Film

	for rows.Next() {

		film := &Film{}
		err := rows.Scan(&film.Film_id, &film.Title, &film.Description, &film.Release_year, &film.Language_id, &film.Rental_duration, &film.Rental_rate, &film.Length, &film.Replacement_cost, &film.Rating, &film.Last_update, &film.Special_features, &film.Fulltext)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, db.ErrNoRecord
			} else {
				return nil, err
			}
		}
		films = append(films, film)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}
