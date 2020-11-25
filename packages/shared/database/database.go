package database

import (
	"fmt"

	"database/sql"
)

var DB *sql.DB

type DBConfig struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Connect() error {
	var err error

	config := &DBConfig{DBHost: "localhost", DBPort: "5432", DBUser: "my_user", DBPass: "password123", DBName: "my_database"}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)

	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	fmt.Println("Connection Opened to Database")
	return nil
}
