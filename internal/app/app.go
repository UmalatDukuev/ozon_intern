package app

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSL      string
}

func Run() {

	config := app.LoadConfig()
	db := repository.NewPostgresDB(config)

	defer db.Close()

	srv := handler.NewGraphQLServer(db)
	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", srv))
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSL:      os.Getenv("DB_SSL"),
	}
}
