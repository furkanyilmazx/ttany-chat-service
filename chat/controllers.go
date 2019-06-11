package chat

import (
	"net/http"
	"ttany-chat-service/common"
	"ttany-chat-service/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func ChatRoomsRetrieve(c *gin.Context) {

	paginationValidator := common.NewPaginationValidator()
	p, _ := paginationValidator.BindQueryString(c, "CreatedAt", "ID")

	log.Info(paginationValidator)

	var rooms = RoomModels{}
	if count, cursors, err := rooms.AllRoomsCursorPaginated(p); err != nil {
		log.Error("Error Occured", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		serializer := RoomSerializer{roomModels: rooms}
		c.JSON(http.StatusOK, gin.H{"_meta_data": gin.H{"count": count, "cursors": cursors}, "data": serializer.ResponseWithArray()})
	}

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
	serializer := RoomSerializer{roomModel: roomModelValitor.roomModel}
	c.JSON(http.StatusOK, serializer.Response())
}

func ChatRoomRetrieve(c *gin.Context) {
}
