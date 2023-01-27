package models

import "time"

// User model
type User struct {
	Username  string `gorm:"primaryKey"`
	Email     string
	Password  string
	CreatedAt time.Time
	LastLogin time.Time
}

type SignupUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupUserResponse struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginUserRequest struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type LoginUserResponse struct {
	AccessToken string    `json:"accessToken"`
	LastLogin   time.Time `json:"lastLogin"`
	User        User      `json:"user"`
}
