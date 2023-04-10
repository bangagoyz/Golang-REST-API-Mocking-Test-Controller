package middleware

import (
	"chapter3_2/database"
	"chapter3_2/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CarAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		carId, err := strconv.Atoi(c.Param("carId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "invalid parameter",
			})
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		Car := models.Car{}

		err = db.Select("user_id").First(&Car, uint(carId)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":    "Data Not Found",
				"messages": "data doesnt exist",
			})
			return
		}

		if Car.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}

		c.Next()
	}
}

func AdminAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userEmail := userData["email"].(string)

		if userEmail != "admin@secondcars.com" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
