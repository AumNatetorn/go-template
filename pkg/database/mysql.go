package database

import (
	"database/sql"
	"fmt"
	"go-template/configs"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(config *configs.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Secrets.DbUsername,
		config.Secrets.DbPassword,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("✅ Successfully connected to MySQL")

	return DB, nil
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("❎ Database connection closed")
	}
}
