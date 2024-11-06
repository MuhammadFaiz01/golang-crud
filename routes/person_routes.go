package routes

import (
	"golang-crud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func PersonRoutes(router *gin.Engine, db *pgxpool.Pool) {
	router.POST("/person", controllers.CreatePerson(db))
	router.GET("/person", controllers.GetPerson(db))
	router.PUT("/person", controllers.UpdatePerson(db))
	router.DELETE("/person", controllers.DeletePerson(db))
}
