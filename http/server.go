package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"stock-bit/config"
	"stock-bit/repository"
)

type Server struct {
	Router *gin.Engine
}

func (server *Server) Run(config config.Configuration) {
	config.Init()
	newRepositories, err := repository.NewRepositories(config)
	if err != nil {
		log.Fatalln(err)
	}
	newRepositories.Seeder()
	newRepositories.AddForeignKey()
	server.Router = gin.Default()
	server.routes(newRepositories)
	log.Fatalln(server.Router.Run(":" + config.Port))
}
