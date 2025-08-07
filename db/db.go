package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnDb() *sql.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := "user=" + dbUser +
		" dbname=" + dbName +
		" password=" + dbPassword +
		" host=" + dbHost +
		" port=" + dbPort +
		" sslmode=" + dbSSLMode

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}
	return db
}
