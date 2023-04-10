package controllers

import (
	"chapter3_2/database"
	"chapter3_2/helpers"
	"chapter3_2/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetCarId(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	Car := models.Car{}
	CarId, _ := strconv.Atoi(c.Param("carId"))

	Car.UserID = userId
	Car.ID = uint(CarId)

	err := db.First(&Car).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Car not found",
		})
		return
	}

	c.JSON(http.StatusOK, Car)
}

func CreateCar(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Car := models.Car{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Car)
	} else {
		c.ShouldBind(&Car)
	}

	Car.UserID = userID

	err := db.Debug().Create(&Car).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Car)

}

func UpdateCar(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Car := models.Car{}

	carId, _ := strconv.Atoi(c.Param("carId"))
	userId := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Car)
	} else {
		c.ShouldBind(&Car)
	}

	Car.UserID = userId
	Car.ID = uint(carId)

	err := db.Model(&Car).Where("id = ?", carId).Updates(models.Car{Title: Car.Title, Brand: Car.Brand, Model: Car.Model, Description: Car.Description}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Car)

}

func DeleteCar(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))
	carId, _ := strconv.Atoi(c.Param("carId"))

	err := db.Where("id = ? AND user_id = ?", carId, userId).Delete(&models.Car{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully",
	})
}

func AdminGetCarById(c *gin.Context) {
	db := database.GetDB()
	carIdParam := c.Param("carId")
	Car := models.Car{}
	CarId, _ := strconv.Atoi(carIdParam)

	Car.ID = uint(CarId)

	err := db.First(&Car).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Car not found",
		})
		return
	}

	c.JSON(http.StatusOK, Car)
}

func AdminDeleteCarById(c *gin.Context) {
	db := database.GetDB()
	carId, _ := strconv.Atoi(c.Param("carId"))

	err := db.Where("id = ? ", carId).Delete(&models.Car{}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully",
	})
}
