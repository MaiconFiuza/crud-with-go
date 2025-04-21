package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func SetupDatabase() *sql.DB {
	if os.Getenv("RUNNING_IN_DOCKER") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Erro ao carregar .env local")
		}
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	dbConnection, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()

	fmt.Println(err)

	return dbConnection
}
