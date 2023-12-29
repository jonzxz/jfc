package models

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/jonzxz/jfc/utils"
	"gorm.io/gorm"
)

type Payment struct {
	ID int `json:"ID" gorm:"primaryKey"`
	// pass in as string, convert to uint64 epoch
	Timestamp   int64   `json:"Timestamp" gorm:"column:TIMESTAMP"`
	Type        string  `json:"Type" gorm:"column:TYPE"`
	Remarks     string  `json:"Remarks" gorm:"column:REMARKS"`
	TotalAmount float32 `json:"TotalAmount" gorm:"column:TOTAL_AMOUNT"`
}

func (Payment) TableName() string {
	return "PAYMENT"
}

func GetAllPaymentsHandler(db *gorm.DB) []Payment {
	payments := []Payment{}
	db.Find(&payments)

	return payments
}

func GetPaymentWrapper(db *gorm.DB, params *url.Values) []Payment {
	var payment []Payment
	if params.Has("id") {
		fmt.Printf("Retrieving payments by id: %v\n", params.Get("id"))
		ids := strings.Split(params.Get("id"), ",")
		payment = getPaymentsByIdHandler(db, ids)
	} else if params.Has("type") && params.Has("month") {
		fmt.Printf("Retrieving payments by month: %v and type: %v\n", params.Get("month"), params.Get("type"))
		payment = getPaymentsByMonthAndTypeHander(db, params.Get("month"), params.Get("type"))
	} else if params.Has("type") {
		fmt.Printf("Retrieving payments by type: %v\n", params.Get("type"))
		payment = getPaymentsByTypeHandler(db, params.Get("type"))
	} else if params.Has("month") {
		fmt.Printf("Retrieving payments by month: %v\n", params.Get("month"))
		payment = getPaymentsByMonthHandler(db, params.Get("month"))
	} else {
		db.Limit(10).Find(&payment)
	}

	return payment

}

func getPaymentsByIdHandler(db *gorm.DB, ids []string) []Payment {
	var payments []Payment
	if len(ids) == 1 {
		var payment Payment
		db.First(&payment, ids[0])
		payments = append(payments, payment)
	} else {
		db.Find(&payments, ids)
	}

	return payments
}

func getPaymentsByMonthAndTypeHander(db *gorm.DB, month string, paymentType string) []Payment {
	var payments []Payment
	epochRange := utils.GetStartEndEpochFromMonth(month)
	db.Where("TIMESTAMP BETWEEN ? AND ? AND TYPE = ?", epochRange["start"], epochRange["end"], paymentType).Find(&payments)
	return payments
}

func getPaymentsByTypeHandler(db *gorm.DB, paymentType string) []Payment {

	var payments []Payment
	db.Where("TYPE = ?", paymentType).Find(&payments)
	return payments
}

func getPaymentsByMonthHandler(db *gorm.DB, month string) []Payment {

	var payments []Payment
	epochRange := utils.GetStartEndEpochFromMonth(month)

	db.Where("TIMESTAMP BETWEEN ? AND ?", epochRange["start"], epochRange["end"]).Find(&payments)

	return payments
}

// Adds a Payment row and creates corresponding rows in PaymentDue
func AddPaymentHandler(db *gorm.DB, payment Payment) {
	payment.Timestamp = time.Now().Unix()
	db.Create(&payment)

	AddPaymentDueFromPaymentHandler(db, &payment)
}
