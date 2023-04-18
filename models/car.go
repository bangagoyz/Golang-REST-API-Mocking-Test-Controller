package models

import (
	"time"
)

type Car struct {
	CarID       string `gorm:"primaryKey;type:varchar(255)"`
	Title       string `gorm:"not null;type:varchar(255);default:null"`
	Brand       string `gorm:"not null;type:varchar(255);default:null"`
	Model       string `gorm:"not null;type:varchar(255);default:null"`
	Description string `gorm:"not null;type:varchar(255);default:null"`
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CarRequest struct {
	Title       string `json:"title" valid:"required~Insert title"`
	Brand       string `json:"brand" valid:"required~insert your car brand!"`
	Model       string `json:"model" valid:"required~insert your car model!"`
	Description string `json:"description" valid:"required~insert your car car description!"`
}

type CarCreateResponse struct {
	CarID       string    `json:"car_id"`
	Title       string    `json:"title"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type CarUpdateResponse struct {
	CarID       string    `json:"car_id"`
	Title       string    `json:"title"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	UpdateAt    time.Time `json:"update_at"`
}

type CarResponse struct {
	CarID       string    `json:"car_id"`
	Title       string    `json:"title"`
	Brand       string    `json:"brand"`
	Model       string    `json:"model"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"update_at"`
}
