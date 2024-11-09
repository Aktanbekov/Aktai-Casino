package routes

import (
	controller "github.com/Aktanbekov/Aktai-Casino/controllers"
	"github.com/gin-gonic/gin"
)


func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup())
	incomingRoutes.POST("user/login", controller.Login())
}