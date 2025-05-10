package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func Connect() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("USER")
	cfg.DBName = os.Getenv("DBNAME")
	cfg.Passwd = os.Getenv("PASSWD")
	cfg.Addr = os.Getenv("ADDR")
	cfg.Net = "tcp"

	fmt.Println(cfg.Addr)

	dns := cfg.FormatDSN()

	db, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("Connection error", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error testing connection", err)
	}

	fmt.Println("Connected to the database successfully")
}

func GetDB() *sql.DB {
	return db
}
