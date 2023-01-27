package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rwiteshbera/Paseto-Go-Implementation/config"
	"github.com/rwiteshbera/Paseto-Go-Implementation/token"
)

type Server struct {
	Config     config.Config
	Router     *gin.Engine
	TokenMaker token.PasetoMaker
}

func CreateServer(config *config.Config) (*Server, error) {
	gin.SetMode(gin.ReleaseMode)
	tokenMaker, err := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)
	if err != nil {
		return nil, errors.Wrap(err, "server package 01")
	}
	server := &Server{
		Config:     *config,
		Router:     gin.Default(),
		TokenMaker: *tokenMaker,
	}
	server.Router.Use(gin.Recovery())
	return server, nil
}

func (server *Server) Start(port string) error {
	return server.Router.Run(":" + port)
}
