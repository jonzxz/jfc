package main

import (
	"database/sql"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"
)

var (
	//db *sql.DB
	cfg = mysql.Config{
		User:   "bizuser",
		Passwd: "secretpassword",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "JFC",
	}
)

type Person struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	TelegramId string `json:"TelegramId"`
}

type Payment struct {
	ID int `json:"ID"`
	// pass in as string, convert to uint64 epoch
	Date    string `json:"Date"`
	Remarks string `json:"Remarks"`
}

type PaymentDue struct {
	ID int `json:"ID"`
	// FK Person.ID
	Payer         int    `json:"Payer"`
	PaymentDue    string `json:"PaymentDue"`
	PayableAmount int    `json:"PayableAmount"`
	Paid          bool   `json:"Paid"`
}

//var people = []Person{
//{ID: 1, Name: "Jonny", TelegramId: "@jonny"},
//{ID: 2, Name: "Sally", TelegramId: "@sally"},
//}

func getPeople(c *gin.Context) {

	people := getAllPersonsHandler()
	c.IndentedJSON(http.StatusOK, people)
}

func getAllPersonsHandler() []Person {
	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatalf("err %v\n", err)
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM PERSON")

	if err != nil {
		log.Fatalf("err %v\n", err)
		return nil
	}

	people := []Person{}
	for results.Next() {
		var t Person
		err = results.Scan(&t.ID, &t.Name, &t.TelegramId)

		if err != nil {
			log.Fatalf("err %v\n", err)
		}

		people = append(people, t)
	}

	return people

}

//func addPeople(c *gin.Context) {
//var newPerson Person

//if err := c.BindJSON(&newPerson); err != nil {
//return
//}

//people = append(people, newPerson)
//c.IndentedJSON(http.StatusCreated, newPerson)

//}

func main() {

	// use os.Getenv instead

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	log.Println("Connected to db!")

	router := gin.Default()
	router.GET("/people", getPeople)
	//router.POST("/add", addPeople)

	router.Run("localhost:8080")

}
