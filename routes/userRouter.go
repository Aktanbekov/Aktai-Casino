package routes

import (
	controller "aktai_casino/controllers"
	"aktai_casino/middleware "
	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRouter *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}