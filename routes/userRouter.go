package routes

import (
	controller "github.com/Aktanbekov/Aktai-Casino/controllers"
	"github.com/Aktanbekov/Aktai-Casino/middleware"
	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRouter *gin.Engine) {
	incomingRouter.Use(middleware.Authenticate())
	incomingRouter.GET("/users", controller.GetUsers())
	incomingRouter.GET("/users/:user_id", controller.GetUser())
}