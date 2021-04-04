package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/teeaaspoon/react-go-login/src/backend/database"
	"github.com/teeaaspoon/react-go-login/src/backend/routes"
	"time"
)

func main() {
	// Connect to Database
	db := database.Connect()
	dbConnection, _ := db.DB()
	defer dbConnection.Close()

	// Set up router with cors
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "My-Xsrf-Token"},
		ExposeHeaders:    []string{"Content-Length", "My-Xsrf-Token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	// pass db into our context so we avoid using global db variable
	router.Use(func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	})

	// setup handler functions for router
	routes.Setup(router)

	router.Run(":8000")
}
