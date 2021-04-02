package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/teeaaspoon/react-go-login/src/backend/routes"
	"time"
)

func main() {
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

	routes.Setup(router)

	router.Run(":8000")
}
