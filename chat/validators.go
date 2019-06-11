package chat

import (
	"strings"
	"ttany-chat-service/utils"

	"github.com/gin-gonic/gin"
)

type RoomModelValidator struct {
	Room struct {
		AdminID      string `json:"admin_id" binding:"required"`
		Name         string `json:"name" binding:"required"`
		Type         string `json:"type" binding:"required"`
		Participants []struct {
			UserID string `json:"user_id" binding:"required"`
		} `json:"participants"`
	} `json:"room"`
	roomModel RoomModel `json:"-"`
}

func (self *RoomModelValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, &self.Room); err != nil {
		return err
	}
	participantsLen := len(self.Room.Participants)
	self.roomModel.AdminID = strings.TrimSpace(self.Room.AdminID)
	self.roomModel.Name = strings.TrimSpace(self.Room.Name)
	self.roomModel.Type = strings.TrimSpace(self.Room.Type)
	self.roomModel.ParticipantModels = make([]ParticipantModel, participantsLen)

	for i := 0; i < participantsLen; i++ {
		self.roomModel.ParticipantModels[i].UserID = strings.TrimSpace(self.Room.Participants[i].UserID)
	}

	return nil
}

func NewRoomModelValidator() RoomModelValidator {
	RoomModelValidator := RoomModelValidator{}
	return RoomModelValidator
}
