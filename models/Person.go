package models

import (
	"database/sql"
	"log"

	"github.com/jonzxz/jfc/constants"
)

type Person struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	TelegramId string `json:"TelegramId"`
}

func GetAllPersonsHandler() []Person {
	db, err := sql.Open("mysql", constants.DB_CONFIGS.FormatDSN())

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
