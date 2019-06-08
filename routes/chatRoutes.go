package routes

import (
	"ttany-chat-service/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func loadChatRoutes(r *gin.RouterGroup, db *gorm.DB) {
	chatRoutes := r.Group("/c")
	chatController := controllers.ChatController{DB: db}
	chatRoutes.GET("/", chatController.GetChatRoomsController)
	chatRoutes.GET("/:room_id", chatController.CreateChatRoomController)
}
