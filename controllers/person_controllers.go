package controllers

import (
	"context"
	"golang-crud/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreatePerson(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var person models.Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec(context.Background(), "INSERT INTO person (full_name, age, birth_date, address) VALUES ($1, $2, $3, $4)",
			person.FullName, person.Age, person.BirthDate, person.Address)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, person)
	}
}

func UpdatePerson(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Query("id"))
		var person models.Person
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec(context.Background(), "UPDATE person SET full_name = $1, age = $2, birth_date = $3, address = $4 WHERE id = $5",
			person.FullName, person.Age, person.BirthDate, person.Address, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
	}
}

func GetPerson(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fullName := c.Query("full_name")
		if fullName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Full name is required"})
			return
		}

		var person models.Person

		err := db.QueryRow(context.Background(),
			"SELECT id, full_name, age, birth_date, address FROM person WHERE full_name = $1", fullName).
			Scan(&person.Id, &person.FullName, &person.Age, &person.BirthDate, &person.Address)

		if err != nil {
			log.Printf("Error querying database: %v", err)

			if err == pgx.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database", "details": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, person)
	}
}

func DeletePerson(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		log.Println("Received ID:", idStr)
		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		commandTag, err := db.Exec(context.Background(), "DELETE FROM person WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if commandTag.RowsAffected() == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
	}
}
