package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "My-Xsrf-Token"}, // allowed headers coming from request
		ExposeHeaders:    []string{"Content-Length", "My-Xsrf-Token"}, // allow the front end to see these headers i send
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	router.Run(":8000")
}
