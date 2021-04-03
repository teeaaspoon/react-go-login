package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func Connect() {
	fmt.Println("CONNECT")

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// sslmode would be enabled in production we'll assume theres secure communication already
	dsn := fmt.Sprintf("host=database user=%s password=%s dbname=%s port=5432 sslmode=disable", dbUser, dbPassword, dbName)

	_, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbError != nil {
		panic("could not connect to the database")
	}

	fmt.Println("CONNECTED TO DATABASE")
}
