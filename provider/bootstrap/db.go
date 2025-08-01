package bootstrap

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("PostgreSQL connected successfully!", os.Getpid())
	return DB
}
