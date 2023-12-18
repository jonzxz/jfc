package main

import (
	"database/sql"
	"fmt"
	"net/http"

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
	//router.GET("/people", getPeople)
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

	router.Run("localhost:8080")

}
