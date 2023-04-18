package main

import (
	"chapter3_2/database"
	"chapter3_2/router"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {

	routers := gin.Default()

	database.StartDB()
	db := database.GetDB()
	router.StartApp(routers, db)

	routers.Run(":" + PORT)
}
