package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName +
		" host=" + dbHost + " port=" + dbPort + " sslmode=" + dbSSLMode

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error obtaining the DB instance:", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	log.Println("Database connection successful")
}
func GetDB() *gorm.DB {
	return DB
}
