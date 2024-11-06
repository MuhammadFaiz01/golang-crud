package main

import (
	"golang-crud/config"
	"golang-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()
	router := gin.Default()
	routes.PersonRoutes(router, db)
	router.Run(":8080")
}