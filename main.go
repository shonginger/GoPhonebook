package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/shonginger/GoPhonebook/Lehem/cmd/api"
	"github.com/shonginger/GoPhonebook/Lehem/config"
	"github.com/shonginger/GoPhonebook/Lehem/db"
)

func main() {
	fmt.Println("server start")

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	InitStorage(db)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8081", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
