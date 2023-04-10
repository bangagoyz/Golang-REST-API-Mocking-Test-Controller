package main

import (
	"chapter3_2/database"
	"chapter3_2/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
