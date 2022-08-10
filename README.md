# tesTask
### Installing the application
Open folder `server` in terminal and `go mod tidy`
### Running the application
Run `docker-compose.yml` and `go run main.go`
### Testing and requests
#### The server handles requests of the type:
1) `localhost:8080/film?id=` getting information about the movie with ID,
2) `localhost:8080/film?title=` getting information about a movie with Title,
3) `localhost:8080/film/create` adding a test movie
4) `localhost:8080/film/list?rating=` getting the list of movies with Rental_rating more than Rating.
#### DB has a table ResponseTimeLog in which all requests and time
#### And if you want to connect to PostgreSQL or Redis
1) Connection string of PostgreSQL DB: `user=postgres password=123456 dbname=dvdrental sslmode=disable`
or `postgres://postgres:123456@localhost:5432/dvdrental?sslmode=disable`
2) Connection string of Redis: url:`localhost:6379` pass:`123456`

