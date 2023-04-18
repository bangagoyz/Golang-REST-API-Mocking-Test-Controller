package controllers

import (
	"chapter3_2/models"
	"chapter3_2/service"

	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CarController struct {
	CarService service.CarService
}

func NewCarController(carService service.CarService) *CarController {
	return &CarController{
		CarService: carService,
	}
}

func (sc *CarController) CreateCar(ctx *gin.Context) {
	var request models.CarRequest
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

	valid, err := valid.ValidateStruct(request)

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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: models.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := sc.CarService.Create(request, userId.(string))
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

	ctx.JSON(http.StatusCreated, models.SuccessResponse{
		Response: models.Meta{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: res,
	})
	return
}

func (sc *CarController) GetAllCar(ctx *gin.Context) {
	AllCar, err := sc.CarService.GetAll()
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
		Data: AllCar,
	})
	return
}

func (sc *CarController) GetOneCar(ctx *gin.Context) {
	CarId := ctx.Param("car_id")
	getCar, err := sc.CarService.GetOne(CarId)

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, models.SuccessResponse{
		Response: models.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: getCar,
	})
}

func (sc *CarController) UpdateCar(ctx *gin.Context) {
	var UpCar models.CarRequest
	CarId := ctx.Param("car_id")

	if err := ctx.ShouldBindJSON(&UpCar); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(UpCar)
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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: models.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := sc.CarService.Update(UpCar, CarId, userId.(string))

	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
			})
			return
		} else if err == models.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				},
				Error: err.Error(),
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
		Data: res,
	})
	return
}

func (sc *CarController) DeleteCar(ctx *gin.Context) {
	CarId := ctx.Param("car_id")
	UserID, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.FailedResponse{
			Response: models.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: models.ErrorInvalidToken.Err,
		})
		return
	}

	err := sc.CarService.Delete(CarId, UserID.(string))
	if err != nil {
		if err == models.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: models.ErrorNotFound.Err,
			})
			return
		} else if err == models.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.FailedResponse{
				Response: models.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
				},
				Error: err.Error(),
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
		Data: "Delete Car success",
	})
	return

}
