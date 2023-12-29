package models

import (
	"github.com/jonzxz/jfc/utils"
	"gorm.io/gorm"
)

type PaymentDue struct {
	ID int `json:"ID" gorm:"primaryKey"`
	// FK Person.ID
	PayerID int `json:"Payer" gorm:"column:PAYER_ID"`
	// FK Payment.ID
	PaymentID      int     `json:"PaymentID" gorm:"column:PAYMENT_ID"`
	PaymentDueDate int64   `json:"PaymentDueDate" gorm:"column:PAYMENT_DUE_TIMESTAMP"`
	PayableAmount  float32 `json:"PayableAmount" gorm:"column:PAYABLE_AMOUNT"`
	Paid           bool    `json:"Paid" gorm:"column:PAID"`
}

func (PaymentDue) TableName() string {
	return "PAYMENT_DUE"
}

func AddPaymentDueFromPaymentHandler(db *gorm.DB, payment *Payment) {

	people := GetAllPersonsHandler(db)
	numOfPayablePax := len(people)
	individualAmountPayable := payment.TotalAmount / float32(numOfPayablePax)
	paymentDueTimestamp := utils.GetLastEpochOfCurrentMonthFromEpoch(payment.Timestamp)

	for _, p := range people {
		paymentDue := PaymentDue{
			PayerID:        p.ID,
			PaymentID:      payment.ID,
			PaymentDueDate: paymentDueTimestamp,
			PayableAmount:  individualAmountPayable,
			Paid:           false,
		}
		db.Create(&paymentDue)
	}
}

func GetAllPaymentDueHandler(db *gorm.DB) []PaymentDue {
	paymentDue := []PaymentDue{}

	db.Find(&paymentDue)

	return paymentDue
}

func UpdatePaymentDuePaidHandler(db *gorm.DB, paymentDue PaymentDue) {
	db.First(&paymentDue)
	paymentDue.Paid = true
	db.Save(&paymentDue)
}
