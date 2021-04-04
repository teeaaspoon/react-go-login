package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teeaaspoon/react-go-login/src/backend/controller"
)

// Setup sets up the routes to their HandlerFunc
func Setup(router *gin.Engine) {
	router.POST("/api/login", controller.Login)
	router.POST("/api/logout", controller.Logout)
}
