package constants

import (
	"github.com/go-sql-driver/mysql"
)

var (
	//db *sql.DB
	cfg = mysql.Config{
		User:   "bizuser",
		Passwd: "secretpassword",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "JFC",
	}
)
