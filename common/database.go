// common/database.go

package common

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDatabase() {
	// Initialize database connection (replace with your connection details)
	var err error
	db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/carpool")
	if err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}
