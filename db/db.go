package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(db:3306)/phonebook_db")
	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
