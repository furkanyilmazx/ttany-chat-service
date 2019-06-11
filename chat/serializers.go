package chat

import (
	"time"

	"github.com/gin-gonic/gin"
)

type RoomSerializer struct {
	roomModel  RoomModel
	roomModels RoomModels
}

type RoomResponse struct {
	RoomID    string    `json:"room_id"`
	AdminID   string    `json:"admin_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`   // 'direct|group'
	Status    string    `json:"status"` // 'closed|fraud' and so on
	CreatedAt time.Time `json:"created_at"`
}

func (self *RoomSerializer) Response() RoomResponse {
	room := RoomResponse{
		RoomID:    self.roomModel.RoomID,
		AdminID:   self.roomModel.AdminID,
		Name:      self.roomModel.Name,
		Type:      self.roomModel.Type,
		CreatedAt: self.roomModel.CreatedAt,
	}
	return room
}

func (self *RoomSerializer) ResponseWithArray() []RoomResponse {
	length := len(self.roomModels)
	roomResponseArray := make([]RoomResponse, length)
	for i := 0; i < len(self.roomModels); i++ {
		self.roomModel = self.roomModels[i]
		roomResponseArray[i] = self.Response()
	}

	return roomResponseArray
}

type ParticipantSerializer struct {
	c *gin.Context
	ParticipantModel
}

type ParticipantResponse struct {
	UserID    string    `json:"user_id" binding:"exists"`
	CreatedAt time.Time `json:"created_at"`
}

func (self *ParticipantSerializer) Response() ParticipantResponse {
	participant := ParticipantResponse{
		UserID:    self.ParticipantModel.UserID,
		CreatedAt: self.ParticipantModel.CreatedAt,
	}
	return participant
}
