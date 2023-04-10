package router

import (
	"chapter3_2/controllers"
	middleware "chapter3_2/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	carRouter := r.Group("/cars")
	{
		carRouter.Use(middleware.Authentication())
		carRouter.POST("/", controllers.CreateCar)
		carRouter.PUT("/:carId", middleware.CarAuthorization(), controllers.UpdateCar)
		carRouter.DELETE("/:carId", middleware.CarAuthorization(), controllers.DeleteCar)
		carRouter.GET("/:carId", middleware.CarAuthorization(), controllers.GetCarId)

		carRouter.PUT("/admin/:carId", middleware.AdminAccess(), controllers.UpdateCar)
		carRouter.DELETE("/admin/:carId", middleware.AdminAccess(), controllers.AdminDeleteCarById)
		carRouter.GET("/admin/:carId", middleware.AdminAccess(), controllers.AdminGetCarById)
	}

	return r
}
