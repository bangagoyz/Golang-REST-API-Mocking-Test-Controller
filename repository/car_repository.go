package repository

import (
	"chapter3_2/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockery --name ICarRepository
type ICarRepository interface {
	Add(newCar models.Car) (models.Car, error)
	Get() ([]models.Car, error)
	GetOne(CarID string) (models.Car, error)
	Update(updateCar models.Car, carId string) (models.Car, error)
	Delete(carID string) error
}

type CarRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{
		db: db,
	}
}

func (sr *CarRepository) Add(newCar models.Car) (models.Car, error) {
	tx := sr.db.Create(&newCar)
	return newCar, tx.Error
}

func (sr *CarRepository) Get() ([]models.Car, error) {
	car := []models.Car{}

	tx := sr.db.Find(&car)
	return car, tx.Error
}

func (sr *CarRepository) GetOne(CarID string) (models.Car, error) {
	FindCar := models.Car{}

	err := sr.db.Debug().Where("car_id = ?", CarID).Take(&FindCar).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Car{}, models.ErrorNotFound
	}
	return FindCar, err
}

func (sr *CarRepository) Update(updateCar models.Car, carId string) (models.Car, error) {
	err := sr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "car_id"},
			{Name: "title"},
			{Name: "updated_at"},
		},
	},
	).
		Where("car_id = ?", carId).Updates(&updateCar)
	return updateCar, err.Error
}

func (sr *CarRepository) Delete(carID string) error {
	delCar := models.Car{}

	tx := sr.db.Delete(&delCar, "social_id = ?", carID)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
