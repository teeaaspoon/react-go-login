package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/teeaaspoon/react-go-login/src/backend"
)


func Setup(router *gin.Engine) {
	router.POST("/api/login", controllers.Login)
}
