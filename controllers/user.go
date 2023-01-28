package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/Paseto-Go-Implementation/middlewares"
	"github.com/rwiteshbera/Paseto-Go-Implementation/server"
	"net/http"
)

func GetUserData(server *server.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Query("username")
		payload, payloadExists := context.Get(middlewares.AuthorizationPayloadKey)
		if !payloadExists {
			context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unable to fetch authorized user's data"})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": username, "payload": payload})
	}
}
