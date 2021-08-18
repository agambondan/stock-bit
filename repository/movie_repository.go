package repository

import (
	"database/sql"
	"fmt"
	"stock-bit/model"
)

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepository{db}
}

type MovieRepository interface {
	Save(movie *model.Movie) (*model.Movie, error)
	SaveMovieDetail(movieDetail *model.MovieDetail) (*model.MovieDetail, error)
	GetAllMovie() ([]model.Movie, error)
	GetMovieDetails() ([]model.MovieDetail, error)
}

func (m *movieRepository) Save(movie *model.Movie) (*model.Movie, error) {
	query := fmt.Sprintf("insert into movies (title, year, imdb_id, type, poster) values ($1, $2, $3, $4, $5) returning id;")
	prepare, err := m.db.Prepare(query)
	if err != nil {
		return movie, err
	}
	err = prepare.QueryRow(&movie.Title, &movie.Year, &movie.Imdbid, &movie.Type, &movie.Poster).Scan(&movie.ID)
	if err != nil {
		return movie, err
	}
	return movie, err
}

func (m *movieRepository) GetAllMovie() ([]model.Movie, error) {
	var movies []model.Movie
	query := fmt.Sprintf("select * from movie")
	prepare, err := m.db.Prepare(query)
	if err != nil {
		return movies, err
	}
	rows, err := prepare.Query()
	for rows.Next() {
		var movie model.Movie
		err = rows.Scan(&movie.ID, &movie.Title, &movie.Year, &movie.Imdbid, &movie.Type, &movie.Poster)
		if err != nil {
			return movies, err
		}
		movies = append(movies, movie)
	}
	return movies, err
}

func (m *movieRepository) SaveMovieDetail(movieDetail *model.MovieDetail) (*model.MovieDetail, error) {
	var id int
	query := fmt.Sprintf("insert into movies_detail (title, year, rated, released, runtime, genre, director, writer, actors, plot, language, country, awards, poster) " +
		"values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) returning id;")
	prepare, err := m.db.Prepare(query)
	if err != nil {
		return movieDetail, err
	}
	err = prepare.QueryRow(&movieDetail.Title, &movieDetail.Year, &movieDetail.Rated, &movieDetail.Released, &movieDetail.Runtime, &movieDetail.Genre, &movieDetail.Director,
		&movieDetail.Writer, &movieDetail.Actors, &movieDetail.Plot, &movieDetail.Language, &movieDetail.Country, &movieDetail.Awards, &movieDetail.Poster).Scan(&id)
	if err != nil {
		return movieDetail, err
	}
	return movieDetail, err
}

func (m *movieRepository) GetMovieDetails() ([]model.MovieDetail, error) {
	var id int
	var movieDetails []model.MovieDetail
	query := fmt.Sprintf("select * from movies_detail")
	prepare, err := m.db.Prepare(query)
	if err != nil {
		return movieDetails, err
	}
	rows, err := prepare.Query()
	if err != nil {
		return movieDetails, err
	}
	for rows.Next() {
		var movieDetail model.MovieDetail
		err := rows.Scan(&id, &movieDetail.Title, &movieDetail.Year, &movieDetail.Rated, &movieDetail.Released, &movieDetail.Runtime, &movieDetail.Genre, &movieDetail.Director,
			&movieDetail.Writer, &movieDetail.Actors, &movieDetail.Plot, &movieDetail.Language, &movieDetail.Country, &movieDetail.Awards, &movieDetail.Poster)
		if err != nil {
			return movieDetails, err
		}
		movieDetails = append(movieDetails, movieDetail)
	}
	return movieDetails, err
}
