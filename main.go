package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/Paseto-Go-Implementation/config"
	"github.com/rwiteshbera/Paseto-Go-Implementation/routes"
	"github.com/rwiteshbera/Paseto-Go-Implementation/server"
	"log"
	"net/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config:", err)
	}

	Server, err := server.CreateServer(config)
	if err != nil {
		log.Fatalln("error:", err)
	}

	if config.SERVER_PORT == "" {
		config.SERVER_PORT = "5001"
	}

	Server.Router.GET("/api", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	routes.AuthenticationRoutes(Server)

	err = Server.Start(config.SERVER_PORT)
	if err != nil {
		log.Fatalln("error:", err)
	}
}
