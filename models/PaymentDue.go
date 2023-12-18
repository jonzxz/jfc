package models

type PaymentDue struct {
	ID int `json:"ID"`
	// FK Person.ID
	Payer         int    `json:"Payer"`
	PaymentDue    string `json:"PaymentDue"`
	PayableAmount int    `json:"PayableAmount"`
	Paid          bool   `json:"Paid"`
}
