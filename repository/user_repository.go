package repository

import (
	"chapter3_2/models"

	"gorm.io/gorm"
)

//go:generate mockery --name IUserRepository
type IUserRepository interface {
	Add(newUser models.User) (models.User, error)
	GetByEmail(email string) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Add(newUser models.User) (models.User, error) {
	tx := ur.db.Create(&newUser)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return newUser, nil
}

func (ur *UserRepository) GetByEmail(email string) (models.User, error) {
	var user models.User
	tx := ur.db.First(&user, "email = ?", email)
	return user, tx.Error
}
