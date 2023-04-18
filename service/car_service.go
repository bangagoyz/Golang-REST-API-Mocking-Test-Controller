package service

import (
	"chapter3_2/helper"
	"chapter3_2/models"
	"chapter3_2/repository"
)

type ICarService interface {
	Create(request models.CarRequest, userID string) (models.CarCreateResponse, error)
	GetAll() ([]models.CarResponse, error)
	Update(updateReq models.CarRequest, CarID string, userID string) (models.CarUpdateResponse, error)
	GetOne(CarID string) (models.CarResponse, error)
	Delete(CarID string, userID string) error
}

type CarService struct {
	CarRepository repository.ICarRepository
}

func NewCarService(carRepository repository.ICarRepository) *CarService {
	return &CarService{
		CarRepository: carRepository,
	}
}

func (ss *CarService) Create(request models.CarRequest, userID string) (models.CarCreateResponse, error) {
	carID := helper.GenerateID()

	NewCar := models.Car{
		CarID:       carID,
		Title:       request.Title,
		Brand:       request.Brand,
		Model:       request.Model,
		Description: request.Description,
		UserID:      userID,
	}

	res, err := ss.CarRepository.Add(NewCar)
	if err != nil {
		if err != models.ErrorNotFound {
			return models.CarCreateResponse{}, err
		}
		return models.CarCreateResponse{}, models.ErrorNotFound
	}

	response := models.CarCreateResponse{
		CarID:       res.CarID,
		Title:       res.Title,
		Brand:       res.Brand,
		Model:       res.Model,
		Description: res.Description,
		UserID:      res.UserID,
		CreatedAt:   res.CreatedAt,
	}
	return response, nil
}

func (ss *CarService) GetAll() ([]models.CarResponse, error) {
	CarRes := []models.CarResponse{}

	res, err := ss.CarRepository.Get()
	if err != nil {
		return []models.CarResponse{}, err
	}
	for _, CarResponse := range res {
		CarRes = append(CarRes, models.CarResponse{
			CarID:       CarResponse.CarID,
			Title:       CarResponse.Title,
			Brand:       CarResponse.Brand,
			Model:       CarResponse.Model,
			Description: CarResponse.Brand,
			UserID:      CarResponse.UserID,
			CreatedAt:   CarResponse.CreatedAt,
			UpdatedAt:   CarResponse.UpdatedAt,
		})
	}

	return CarRes, nil
}

func (ss *CarService) Update(updateReq models.CarRequest, CarID string, userID string) (models.CarUpdateResponse, error) {
	getId, err := ss.CarRepository.GetOne(CarID)
	if err != nil {
		if err != models.ErrorNotFound {
			return models.CarUpdateResponse{}, err
		}
		return models.CarUpdateResponse{}, models.ErrorNotFound
	}

	if getId.UserID != userID {
		return models.CarUpdateResponse{}, models.ErrorForbiddenAccess
	}

	CarUpdate := models.Car{
		Title:       updateReq.Title,
		Brand:       updateReq.Brand,
		Model:       updateReq.Model,
		Description: updateReq.Description,
	}

	res, err := ss.CarRepository.Update(CarUpdate, CarID)

	if err != nil {
		return models.CarUpdateResponse{}, err
	}

	return models.CarUpdateResponse{
		CarID:       res.CarID,
		Title:       res.Title,
		Brand:       res.Brand,
		Model:       res.Model,
		Description: res.Description,
		UserID:      userID,
		UpdateAt:    res.UpdatedAt,
	}, nil
}

func (ss *CarService) GetOne(CarID string) (models.CarResponse, error) {
	getOne, err := ss.CarRepository.GetOne(CarID)

	if err != nil {
		if err != models.ErrorNotFound {
			return models.CarResponse{}, err
		}
		return models.CarResponse{}, models.ErrorNotFound
	}

	return models.CarResponse{
		CarID:     getOne.CarID,
		Title:     getOne.Title,
		Brand:     getOne.Brand,
		Model:     getOne.Model,
		UserID:    getOne.UserID,
		CreatedAt: getOne.CreatedAt,
		UpdatedAt: getOne.UpdatedAt,
	}, nil
}

func (ss *CarService) Delete(CarID string, userID string) error {
	getCarId, err := ss.CarRepository.GetOne(CarID)

	if err != nil {
		if err != models.ErrorNotFound {
			return err
		}
		return models.ErrorNotFound
	}

	if getCarId.UserID != userID {
		return models.ErrorForbiddenAccess
	}

	err = ss.CarRepository.Delete(CarID)

	if err != nil {
		return err
	}

	return nil
}
