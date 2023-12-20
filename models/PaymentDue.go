package models

type PaymentDue struct {
	ID int `json:"ID"`
	// FK Person.ID
	PayerID        int     `json:"Payer"`
	PaymentID      int     `json:"PaymentId"`
	PaymentDueDate string  `json:"PaymentDueDate"`
	PayableAmount  float32 `json:"PayableAmount"`
	Paid           bool    `json:"Paid"`
}
