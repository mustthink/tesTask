package data

import "time"

type Film struct {
	Film_id          int       `json:"film_id"`
	Title            string    `json:"title" `
	Description      string    `json:"description"`
	Realease_year    int       `json:"realease_year"`
	Language_id      int       `json:"language_id"`
	Rental_duration  int       `json:"rental_duration"`
	Rental_rate      int       `json:"rental_rate"`
	Length           int       `json:"length"`
	Replacement_cost int       `json:"replacement_cost"`
	Rating           int       `json:"rating"`
	Last_update      time.Time `json:"last_update"`
	Special_features string    `json:"special_features"`
	Fulltext         string    `json:"fulltext"`
}

func (m *Service) InsertFilm(film *Film) error {

	stmt := `insert into Film (Title, Description, Realease_year, Language_id, Rental_duration, Rental_rate, Length, Replacement_cost, Rating, Last_update, Special_features, Fulltext) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	_, err := m.DB.Exec(stmt, film.Title, film.Description, film.Realease_year, film.Language_id, film.Rental_duration, film.Rental_rate, film.Length, film.Replacement_cost, film.Rating, film.Last_update, film.Special_features, film.Fulltext)
	if err != nil {
		return err
	}

	return nil
}
