package chat

import (
	"net/http"
	"ttany-chat-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func ChatRoomsRetrieve(c *gin.Context) {
	var rooms = RoomModels{}
	if err := rooms.AllRooms(); err != nil {
		log.Error("Error Occured", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Debug(rooms)
	c.JSON(http.StatusOK, &rooms)
}

func ChatRoomCreate(c *gin.Context) {
	roomModelValitor := NewRoomModelValidator()
	if err := roomModelValitor.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.NewValidatorError(err))
		return
	}
	roomModelValitor.roomModel.RoomID = uuid.New().String()
	for idx, _ := range roomModelValitor.roomModel.ParticipantModels {
		roomModelValitor.roomModel.ParticipantModels[idx].RoomID = roomModelValitor.roomModel.RoomID
	}
	roomModelValitor.roomModel.CreateRoom()
	log.Info(roomModelValitor.roomModel)
	serializer := RoomSerializer{c, roomModelValitor.roomModel}
	c.JSON(http.StatusOK, serializer.Response())
}

func ChatRoomRetrieve(c *gin.Context) {
}
