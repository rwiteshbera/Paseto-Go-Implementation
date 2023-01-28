package routes

import (
	"github.com/rwiteshbera/Paseto-Go-Implementation/controllers"
	"github.com/rwiteshbera/Paseto-Go-Implementation/middlewares"
	"github.com/rwiteshbera/Paseto-Go-Implementation/server"
)

func UserRoutes(server *server.Server) {
	server.Router.Use(middlewares.AuthMiddleware(server))
	server.Router.GET("/user", controllers.GetUserData(server))
}
