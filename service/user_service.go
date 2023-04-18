package service

import (
	"chapter3_2/helper"
	"chapter3_2/models"
	"chapter3_2/repository"
)

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Register(userRegisterRequest models.UserRegisterRequest) (*models.UserRegisterResponse, error) {
	userId := helper.GenerateID()

	hashPassword, err := helper.Hash(userRegisterRequest.Password)
	if err != nil {
		return &models.UserRegisterResponse{}, err
	}

	user := models.User{
		UserID:   userId,
		FullName: userRegisterRequest.FullName,
		Email:    userRegisterRequest.Email,
		Password: hashPassword,
	}

	res, err := us.UserRepository.Add(user)

	if err != nil {
		return &models.UserRegisterResponse{}, err
	}

	return &models.UserRegisterResponse{
		UserID:    res.UserID,
		FullName:  res.FullName,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (us *UserService) Login(userLoginRequest models.UserLoginRequest) (models.UserLoginResponse, error) {

	user, err := us.UserRepository.GetByEmail(userLoginRequest.Email)
	if err != nil {
		return models.UserLoginResponse{}, models.ErrorInvalidEmailOrPassword
	}

	if !helper.IsHashValid(user.Password, userLoginRequest.Password) {
		return models.UserLoginResponse{}, models.ErrorInvalidEmailOrPassword
	}

	token, err := helper.GenerateAccessToken(user.UserID, user.Email)
	if err != nil {
		return models.UserLoginResponse{}, models.ErrorInvalidToken
	}

	return models.UserLoginResponse{
		Token: token,
	}, nil

}
