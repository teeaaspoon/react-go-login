package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/teeaaspoon/react-go-login/src/backend/database"
	"github.com/teeaaspoon/react-go-login/src/backend/routes"
)

func main() {
	// Connect to Database
	db := database.Connect()
	dbConnection, _ := db.DB()
	defer dbConnection.Close()

	// Set up router with cors
	router := gin.Default()
	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://localhost"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Csrf-Token"},
		ExposeHeaders:    []string{"Content-Length", "Csrf-Token"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	// pass db into our context so we avoid using global db variable
	router.Use(func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	})

	// setup handler functions for router
	routes.Setup(router)

	router.RunTLS(":8000", "./ssl/ca-cert.pem", "./ssl/ca-key.pem")
	//router.Run(":8000")
}
