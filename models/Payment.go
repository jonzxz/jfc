package models

type Payment struct {
	ID int `json:"ID"`
	// pass in as string, convert to uint64 epoch
	Date        string  `json:"Date"`
	Remarks     string  `json:"Remarks"`
	TotalAmount float32 `json:"TotalAmount"`
}
