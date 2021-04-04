package database

import (
	"github.com/teeaaspoon/react-go-login/src/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	//dbUser := os.Getenv("POSTGRES_USER")
	//dbPassword := os.Getenv("POSTGRES_PASSWORD")
	//dbName := os.Getenv("POSTGRES_DB")

	// sslmode would be enabled in production we'll assume theres secure communication already
	//dsn := fmt.Sprintf("host=database user=%s password=%s dbname=%s port=5432 sslmode=disable", dbUser, dbPassword, dbName)
	dsn := "host=localhost user=tommypoon dbname=goauth port=5432 sslmode=disable"

	db, dbError := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbError != nil {
		panic("could not connect to the database")
	}

	migrateAndInsertData(db)

	return db
}

// migrateAndInsertData will migrate the users table into our database
// it will also insert one user into the users table for default data
func migrateAndInsertData(db *gorm.DB) {
	migrationError := db.AutoMigrate(&models.User{})

	if migrationError != nil {
		panic("could not migrate database")
	}

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)

	user := models.User{
		Email:    "email@email.com",
		Password: password,
	}

	db.FirstOrCreate(&user)
}
