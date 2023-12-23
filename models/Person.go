package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Person struct {
	ID         int    `json:"ID" gorm:"primaryKey"`
	Name       string `json:"Name" gorm:"column:NAME"`
	TelegramId string `json:"TelegramId" gorm:"column:TELEGRAM_ID"`
}

func (Person) TableName() string {
	return "PERSON"
}

func GetAllPersonsHandler(db *gorm.DB) []Person {
	person := []Person{}
	db.Find(&person)

	return person

}

func AddPersonHandler(db *gorm.DB, person Person) {
	db.Create(&person)

	fmt.Printf("Created user: %v\n", person)

}
