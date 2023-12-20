package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/jonzxz/jfc/constants"
	"github.com/jonzxz/jfc/models"
)

var (
	db *sql.DB
)

func main() {

	db, err := sql.Open("mysql", constants.DB_CONFIGS.FormatDSN())
	fmt.Printf("%p\n", db)

	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Connected to db!")

	router := gin.Default()

	// model.Person
	router.GET("/people", func(c *gin.Context) {
		people := models.GetAllPersonsHandler(db)
		c.IndentedJSON(http.StatusOK, people)

	})
	router.POST("/add", func(c *gin.Context) {
		var newPerson models.Person
		if err := c.BindJSON(&newPerson); err != nil {
			log.Fatalf("%v\n", err)
		}
		models.AddPersonHandler(db, newPerson)
		c.IndentedJSON(http.StatusCreated, newPerson)
	})

	// model.Payment
	router.GET("/payments", func(c *gin.Context) {
		queryParams := c.Request.URL.Query()

		payments := []models.Payment{}

		if queryParams.Get("id") != "" {
			id, err := strconv.Atoi(queryParams.Get("id"))
			if err != nil {
				log.Fatalf("err %v\n", err)
			}
			payments = models.GetSpecificPaymentByIdHandler(db, id)

		} else if queryParams.Get("type") != "" {
			payments = models.GetPaymentsByTypeHandler(db, queryParams.Get("type"))
		} else {
			payments = models.GetAllPaymentsHandler(db)

		}
		c.IndentedJSON(http.StatusOK, payments)
	})

	router.Run("localhost:8080")

}
