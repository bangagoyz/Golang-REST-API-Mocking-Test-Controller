package controllers

import (
	"chapter3_2/models"
	"chapter3_2/service"
	"net/http"

	Valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) *UserController {
	return &UserController{
		UserService: UserService,
	}
}

func (uc *UserController) Register(ctx *gin.Context) {
	var request models.UserRegisterRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := Valid.ValidateStruct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Register(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Response: models.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
	return

}

func (uc *UserController) Login(ctx *gin.Context) {
	var request models.UserLoginRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := Valid.ValidateStruct(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response, err := uc.UserService.Login(request)
	if err != nil {
		if err == models.ErrorInvalidEmailOrPassword {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
				},
				Error: err.Error(),
			})
			return
		} else if err == models.ErrorInvalidToken {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
				},
				Error: models.ErrorInvalidToken.Err,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Response: models.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: response,
	})
}
