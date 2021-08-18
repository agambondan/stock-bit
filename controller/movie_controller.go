package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"stock-bit/model"
	"stock-bit/repository"
	"stock-bit/service"
	"stock-bit/utils"
)

type movieController struct {
	call  service.CallService
	movie service.MovieService
}

type MovieController interface {
	SaveMovie(c *gin.Context)
	SaveMovieDetail(c *gin.Context)
	GetMovies(c *gin.Context)
	GetMovieDetail(c *gin.Context)
	GetMoviesOMDBAPI(c *gin.Context)
	GetMovieDetailOMDBAPI(c *gin.Context)
}

func NewMovieController(repo *repository.Repositories) MovieController {
	newCallService := service.NewCallService()
	newMovieService := service.NewMovieService(repo.Movie)
	return &movieController{call: newCallService, movie: newMovieService}
}

func (m *movieController) SaveMovie(c *gin.Context) {
	var movie model.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	movieCreate, err := m.movie.SaveMovie(&movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, movieCreate)
}

func (m *movieController) SaveMovieDetail(c *gin.Context) {
	var movieDetail model.MovieDetail
	if err := c.ShouldBindJSON(&movieDetail); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	movieDetailCreate, err := m.movie.SaveMovieDetail(&movieDetail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, movieDetailCreate)
}

func (m *movieController) GetMovieDetail(c *gin.Context) {
	movieDetails, err := m.movie.GetMovieDetail()
	if err != nil {
		fmt.Println(movieDetails, err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, movieDetails)
}

func (m *movieController) GetMovies(c *gin.Context) {
}

func (m *movieController) GetMoviesOMDBAPI(c *gin.Context) {
	var movies model.MovieSearch
	pageParam := c.Query("page")
	searchParam := c.Query("s")
	url := fmt.Sprintf("%s/?apikey=%s&s=%s&%s", os.Getenv("URL_OMDAPI"), os.Getenv("API_KEY_OMDBAPI"), searchParam, pageParam)
	response, err := m.call.Request().Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(response.Body(), &movies)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var movie model.Movie
	go utils.WriteFile(fmt.Sprintf("%s\n", string(response.Body())))
	go func() {
		for i := 0; i < len(movies.Search); i++ {
			movie.Title = movies.Search[i].Title
			movie.Poster = movies.Search[i].Poster
			movie.Type = movies.Search[i].Type
			movie.Imdbid = movies.Search[i].Imdbid
			movie.Year = movies.Search[i].Year
			m.call.Request().SetBody(&movie).Post("http://localhost:5000/movies")
		}
	}()
	c.JSON(http.StatusOK, movies)
}

func (m *movieController) GetMovieDetailOMDBAPI(c *gin.Context) {
	var MovieDetail model.MovieDetail
	idParam := c.Param("id")
	url := fmt.Sprintf("%s/?apikey=%s&i=%s", os.Getenv("URL_OMDAPI"), os.Getenv("API_KEY_OMDBAPI"), idParam)
	response, err := m.call.Request().Get(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = json.Unmarshal(response.Body(), &MovieDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	go utils.WriteFile(fmt.Sprintf("%s\n", string(response.Body())))
	go func() {
		m.call.Request().SetBody(&MovieDetail).Post("http://localhost:5000/movies/detail")
	}()
	c.JSON(http.StatusOK, MovieDetail)
}
