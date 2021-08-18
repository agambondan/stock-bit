package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stock-bit/controller"
	"stock-bit/repository"
)

func (server *Server) routes(repositories *repository.Repositories) {
	routes := server.Router

	newMovieController := controller.NewMovieController(repositories)
	//Home Routing
	routes.GET("/", Home)

	routes.POST("/movies", newMovieController.SaveMovie)
	routes.GET("/movies", newMovieController.GetMovies)
	routes.POST("/movies/detail", newMovieController.SaveMovieDetail)
	routes.GET("/movies/detail", newMovieController.GetMovieDetail)

	// call api from omdb and hooks save movies
	routes.GET("/movies/omdb", newMovieController.GetMoviesOMDBAPI)
	routes.GET("/movies/omdb/:id", newMovieController.GetMovieDetailOMDBAPI)
}

func Home(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
