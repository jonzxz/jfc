package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/jonzxz/jfc/constants"
	"github.com/jonzxz/jfc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *sql.DB
)

func main() {

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:                       constants.DB_CONFIGS.FormatDSN(),
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		}), &gorm.Config{})

	fmt.Printf("%p\n", db)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to db!")

	router := gin.Default()

	// model.Person
	router.GET("/people/list", func(c *gin.Context) {
		people := models.GetAllPersonsHandler(db)
		c.IndentedJSON(http.StatusOK, people)

	})

	router.POST("/people/add", func(c *gin.Context) {
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

		payments := models.GetPaymentWrapper(db, &queryParams)

		//}
		c.IndentedJSON(http.StatusOK, payments)
	})

	router.POST("/payments/add", func(c *gin.Context) {
		var newPayment models.Payment
		if err := c.BindJSON(&newPayment); err != nil {
			log.Fatalf("%v\n", err)
		}
		models.AddPaymentHandler(db, newPayment)
		c.IndentedJSON(http.StatusCreated, newPayment)
	})

	// model.PaymentDue
	router.GET("/due", func(c *gin.Context) {
		paymentDue := models.GetAllPaymentDueHandler(db)
		c.IndentedJSON(http.StatusOK, paymentDue)
	})

	router.POST("/due/pay", func(c *gin.Context) {
		var paymentDue models.PaymentDue
		if err := c.BindJSON(&paymentDue); err != nil {
			log.Fatalf("%v\n", err)
		}
		models.UpdatePaymentDuePaidHandler(db, paymentDue)
		c.IndentedJSON(http.StatusOK, paymentDue)
	})

	router.Run("localhost:8080")

}
