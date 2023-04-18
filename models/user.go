package models

import (
	"time"
)

type User struct {
	UserID    string `gorm:"primaryKey;type:varchar(255)"`
	FullName  string `gorm:"unique;not null;type:varchar(255);default:null"`
	Email     string `gorm:"unique;not null;type:varchar(255);default:null"`
	Password  string `gorm:"not null;type:varchar(255)"`
	Cars      []Car
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterRequest struct {
	FullName string `json:"fullname" valid:"required~please input your name"`
	Email    string `json:"email" valid:"required~please input your email,email"`
	Password string `json:"password" form:"password" valid:"required,minstringlength(6)"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required~Username is required"`
	Password string `json:"password" validate:"required~Password is required"`
}

type UserRegisterResponse struct {
	UserID    string    `json:"id"`
	FullName  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
