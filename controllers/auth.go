package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rwiteshbera/Paseto-Go-Implementation/db"
	"github.com/rwiteshbera/Paseto-Go-Implementation/models"
	"github.com/rwiteshbera/Paseto-Go-Implementation/server"
	"github.com/rwiteshbera/Paseto-Go-Implementation/utils"
	"net/http"
	"time"
)

func SignUp(server *server.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var signUpRequest models.SignupUserRequest
		var user models.User
		var signUpResponse models.SignupUserResponse

		err := context.ShouldBindJSON(&signUpRequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		database, err := db.ConnectDB(server.Config)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		database.Migrator().AutoMigrate(&models.User{})

		rows := database.Find(&user, "username = ? OR email = ?", signUpRequest.Username, signUpRequest.Email).RowsAffected
		if rows > 0 {
			context.JSON(http.StatusAccepted, gin.H{"message": "user already exists"})
			return
		}

		// Hash The Password
		password, err := utils.HashPassword(signUpRequest.Password)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		createdAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user = models.User{
			Username:  signUpRequest.Username,
			Email:     signUpRequest.Email,
			Password:  password,
			CreatedAt: createdAt,
			LastLogin: createdAt,
		}

		database.Create(&user)

		signUpResponse = models.SignupUserResponse{
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}

		context.JSON(http.StatusOK, gin.H{"message": signUpResponse})
	}
}

func Login(server *server.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		var loginRequest models.LoginUserRequest
		var user models.User
		var loginResponse models.LoginUserResponse

		err := context.ShouldBindJSON(&loginRequest)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		database, err := db.ConnectDB(server.Config)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		database.Migrator().AutoMigrate(&models.User{})

		rows := database.Find(&user, "username = ? OR email = ?", loginRequest.UsernameOrEmail, loginRequest.UsernameOrEmail).RowsAffected
		if rows == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"error": "no user found"})
			return
		}

		if !utils.VerifyPassword(user.Password, loginRequest.Password) {
			context.JSON(http.StatusBadRequest, gin.H{"error": "incorrect credentials"})
			return
		}

		duration, err := time.ParseDuration(server.Config.ACCESS_TOKEN_DURATION)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		accessToken, err := server.TokenMaker.CreateToken(user.Username, duration)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		loginResponse = models.LoginUserResponse{
			AccessToken: accessToken,
			LastLogin:   user.LastLogin,
			User:        user,
		}

		context.JSON(http.StatusOK, gin.H{"message": loginResponse})
	}
}
