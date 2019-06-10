package chat

import (
	"github.com/gin-gonic/gin"
)

func ChatRoutes(r *gin.RouterGroup) {
	chatRoutes := r.Group("/chat")
	{
		chatRoutes.GET("", ChatRoomsRetrieve)
		chatRoutes.POST("", ChatRoomCreate)
		chatRoutes.GET(":room_id", ChatRoomRetrieve)
	}
}
