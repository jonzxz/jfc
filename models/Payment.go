package models

import (
	"database/sql"
	"log"
	"strconv"
	"time"
)

type Payment struct {
	ID int `json:"ID"`
	// pass in as string, convert to uint64 epoch
	Timestamp   int64   `json:"Timestamp"`
	Type        string  `json:"Type"`
	Remarks     string  `json:"Remarks"`
	TotalAmount float32 `json:"TotalAmount"`
}

func GetAllPaymentsHandler(db *sql.DB) []Payment {
	results, err := db.Query("SELECT * FROM PAYMENT")

	if err != nil {
		log.Fatalf("err %v\n", err)
		return nil
	}

	payments := []Payment{}

	for results.Next() {
		var t Payment
		err = results.Scan(&t.ID, &t.Timestamp, &t.Type, &t.Remarks, &t.TotalAmount)
		payments = append(payments, t)
	}

	return payments

}

func GetSpecificPaymentByIdHandler(db *sql.DB, id int) []Payment {
	row := db.QueryRow("SELECT * FROM PAYMENT WHERE ID = ?", id)

	var t Payment

	if err := row.Scan(&t.ID, &t.Timestamp, &t.Type, &t.Remarks, &t.TotalAmount); err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No such Payment by ID %d", id)
			return nil
		}
		log.Fatalf("err %v", err)
	}
	return []Payment{t}
}

func GetPaymentsByTypeHandler(db *sql.DB, paymentType string) []Payment {
	results, err := db.Query("SELECT * FROM PAYMENT WHERE TYPE = ?", paymentType)

	if err != nil {
		log.Fatalf("err %v\n", err)
		return nil
	}

	payments := []Payment{}

	for results.Next() {
		var t Payment
		err = results.Scan(&t.ID, &t.Timestamp, &t.Type, &t.Remarks, &t.TotalAmount)
		payments = append(payments, t)
	}

	return payments
}

func GetPaymentsByMonthHandler(db *sql.DB, month string) []Payment {

	epochRange := getStartEndEpochFromMonth(month)

	results, err := db.Query("SELECT * FROM PAYMENT WHERE TIMESTAMP >= ? AND TIMESTAMP <= ?", epochRange["start"], epochRange["end"])

	if err != nil {
		log.Fatalf("err %v\n", err)
		return nil
	}

	payments := []Payment{}

	for results.Next() {
		var t Payment
		err = results.Scan(&t.ID, &t.Timestamp, &t.Type, &t.Remarks, &t.TotalAmount)
		payments = append(payments, t)
	}

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
