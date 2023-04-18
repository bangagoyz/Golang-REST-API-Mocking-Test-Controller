package router

import (
	"chapter3_2/controllers"
	"chapter3_2/middlewares"
	"chapter3_2/repository"
	"chapter3_2/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartApp(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(*userService)

	CarRepository := repository.NewCarRepository(db)
	CarService := service.NewCarService(CarRepository)
	CarController := controllers.NewCarController(*CarService)

	base := router.Group("/sellcar")
	{
		user := base.Group("/user")
		{
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
		}
		withAuth := base.Group("/cars", middlewares.AuthMiddleware)
		{
			withAuth.POST("/create", CarController.CreateCar)
			withAuth.GET("/get/all", CarController.GetAllCar)
			withAuth.GET("/get/:car_id", CarController.GetOneCar)
			withAuth.PUT("/update/:car_id", CarController.UpdateCar)
			withAuth.DELETE("/delete/:car_id", CarController.DeleteCar)
		}
	}
}
