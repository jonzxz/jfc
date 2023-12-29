package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID                   int    `json:"ID" gorm:"primaryKey"`
	Name                 string `json:"Name" gorm:"column:NAME"`
	TelegramId           string `json:"TelegramId" gorm:"column:TELEGRAM_ID"`
	CreatedTimestamp     int64  `gorm:"column:CREATED_TS"`
	LastUpdatedTimestamp int64  `gorm:"column:LAST_UPDATED_TS"`
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
	person.CreatedTimestamp = time.Now().Unix()
	db.Create(&person)

	fmt.Printf("Created user: %v\n", person)

}
