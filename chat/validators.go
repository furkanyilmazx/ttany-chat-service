package chat

import (
	"ttany-chat-service/utils"

	"github.com/gin-gonic/gin"
)

type RoomModelValidator struct {
	Room struct {
		AdminID      string `json:"admin_id" binding:"exists"`
		Name         string `json:"name" binding:"exists"`
		Type         string `json:"type" binding:"exists"`
		Participants []struct {
			UserID string `json:"user_id" binding:"exists"`
		} `json:"participants"`
	} `json:"room"`
	roomModel RoomModel `json:"-"`
}

func (self *RoomModelValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, &self.Room); err != nil {
		return err
	}
	self.roomModel.AdminID = self.Room.AdminID
	self.roomModel.Name = self.Room.Name
	self.roomModel.Type = self.Room.Type
	self.roomModel.ParticipantModels = make([]ParticipantModel, len(self.Room.Participants))
	for idx, _ := range self.Room.Participants {
		self.roomModel.ParticipantModels[idx].UserID = self.Room.Participants[idx].UserID
	}

	return nil
}

func NewRoomModelValidator() RoomModelValidator {
	RoomModelValidator := RoomModelValidator{}
	return RoomModelValidator
}
