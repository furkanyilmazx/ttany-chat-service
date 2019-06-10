package chat

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type RoomSerializer struct {
	c *gin.Context
	RoomModel
}

type RoomResponse struct {
	RoomID       string                `json:"room_id"`
	AdminID      string                `json:"admin_id"`
	Name         string                `json:"name"`
	Type         string                `json:"type"`   // 'direct|group'
	Status       string                `json:"status"` // 'closed|fraud' and so on
	CreatedAt    time.Time             `json:"created_at"`
	Participants []ParticipantResponse `json:"participants"`
}

func (self *RoomSerializer) Response() RoomResponse {
	log.Info(self)
	participants := make([]ParticipantResponse, len(self.ParticipantModels))

	for idx, _ := range participants {
		participants[idx].UserID = self.RoomModel.ParticipantModels[idx].UserID
		participants[idx].CreatedAt = self.RoomModel.ParticipantModels[idx].CreatedAt
	}
	room := RoomResponse{
		RoomID:       self.RoomModel.RoomID,
		AdminID:      self.RoomModel.AdminID,
		Name:         self.RoomModel.Name,
		Type:         self.RoomModel.Type,
		CreatedAt:    self.RoomModel.CreatedAt,
		Participants: participants,
	}

	return room
}

type ParticipantResponse struct {
	UserID    string    `json:"user_id" binding:"exists"`
	CreatedAt time.Time `json:"created_at"`
}
