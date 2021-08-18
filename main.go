package main

import (
	"github.com/joho/godotenv"
	"stock-bit/config"
	"stock-bit/http"
	"log"
)

var (
	server      http.Server
	configure   config.Configuration
	pathFileEnv = ".env"
)

func init() {
	if err := godotenv.Load(pathFileEnv); err != nil {
		log.Println("no env gotten")
	}
}

func main() {
	server.Run(configure)
}