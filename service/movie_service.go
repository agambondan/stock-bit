package service

import (
	"stock-bit/model"
	"stock-bit/repository"
)

type movieService struct {
	movie repository.MovieRepository
}

// NewMovieService implements the MovieService Interface
func NewMovieService(movieRepository repository.MovieRepository) MovieService {
	return &movieService{movie: movieRepository}
}

type MovieService interface {
	SaveMovie(movie *model.Movie) (*model.Movie, error)
	GetAllMovie() ([]model.Movie, error)
	SaveMovieDetail(movieDetail *model.MovieDetail) (*model.MovieDetail, error)
	GetMovieDetail() ([]model.MovieDetail, error)
}

func (m *movieService) SaveMovie(movie *model.Movie) (*model.Movie, error) {
	return m.movie.Save(movie)
}

func (m *movieService) GetAllMovie() ([]model.Movie, error) {
	return m.movie.GetAllMovie()
}

func (m *movieService) SaveMovieDetail(movieDetail *model.MovieDetail) (*model.MovieDetail, error) {
	return m.movie.SaveMovieDetail(movieDetail)
}

func (m *movieService) GetMovieDetail() ([]model.MovieDetail, error) {
	return m.movie.GetMovieDetails()
}
