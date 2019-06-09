package routes

import (
	"ttany-chat-service/controllers"

	"github.com/gin-gonic/gin"
)

func loadChatRoutes(r *gin.RouterGroup) {
	chatRoutes := r.Group("/c")
	chatRoutes.GET("", controllers.GetChatRooms)
	chatRoutes.POST("", controllers.CreateChatRoom)
	chatRoutes.GET(":room_id", controllers.GetChatRoomWithRoomID)
}
