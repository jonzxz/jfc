package models

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

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
	epochRange := getStartEndEpochFromMonth(month)
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
	epochRange := getStartEndEpochFromMonth(month)

	db.Where("TIMESTAMP BETWEEN ? AND ?", epochRange["start"], epochRange["end"]).Find(&payments)

	return payments
}

func getStartEndEpochFromMonth(month string) map[string]int64 {
	now := time.Now()
	currentYear, _, _ := now.Date()
	monthInt, _ := strconv.Atoi(month)
	monthTime := time.Month(monthInt)
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, monthTime, 1, 0, 0, 0, 0, currentLocation)
	// dirty add to make 23:59:59
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1).UTC().Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)

	startUnix := firstOfMonth.Unix()
	endUnix := lastOfMonth.Unix()

	epochs := make(map[string]int64)

	epochs["start"] = startUnix
	epochs["end"] = endUnix

	return epochs

}

func AddPaymentHandler(db *gorm.DB, payment Payment) {
	payment.Timestamp = time.Now().Unix()
	db.Create(&payment)

	fmt.Printf("Created payment with %v\n", payment)
}
