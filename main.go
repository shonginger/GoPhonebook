package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/shonginger/GoPhonebook/Lehem/cmd/api"
	"github.com/shonginger/GoPhonebook/Lehem/config"
	"github.com/shonginger/GoPhonebook/Lehem/db"
)

func main() {
	fmt.Println("server start")

	db, err := db.NewMySQLStorage(mysqlCfg.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	InitStorage(db)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	time.Sleep(5 * time.Second)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
