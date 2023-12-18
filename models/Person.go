package models

import (
	"database/sql"
	"log"
)

type Person struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	TelegramId string `json:"TelegramId"`
}

func GetAllPersonsHandler(db *sql.DB) []Person {

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

func AddPersonHandler(db *sql.DB, person Person) {
	insert, err := db.Query(
		"INSERT INTO PERSON (NAME, TELEGRAM_ID) VALUES (?,?)",
		person.Name, person.TelegramId)

	if err != nil {
		log.Fatalf("err %v\n", err)
	}
	defer insert.Close()
}
