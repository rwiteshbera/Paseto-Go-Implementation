package routes

import (
	"github.com/rwiteshbera/Paseto-Go-Implementation/controllers"
	"github.com/rwiteshbera/Paseto-Go-Implementation/server"
)

func AuthenticationRoutes(server *server.Server) {
	server.Router.POST("/signup", controllers.SignUp(server))
	server.Router.POST("/login", controllers.Login(server))
}
