package routes

import (
	"ttany-chat-service/controllers"

	"github.com/gin-gonic/gin"
)

func loadChatRoutes(r *gin.Engine) {
	r.GET("/c", controllers.AllChatsController)
}
