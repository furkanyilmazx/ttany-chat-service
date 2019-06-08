package controllers

import (
	"ttany-chat-service/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type ChatController struct {
	DB *gorm.DB
}

func (ch *ChatController) GetChatRoomsController(c *gin.Context) {
	var rooms = []models.Room{}
	if ch.DB.Preload("Participants").Find(&rooms).RecordNotFound() {
		log.Error("NOT FOUND")
		c.AbortWithStatus(404)
	}
	log.Debug(rooms)
	c.JSON(200, &rooms)
}

func (ch *ChatController) CreateChatRoomController(c *gin.Context) {
	roomID := c.Param("room_id")
	var room = models.Room{}
	var participants = []models.Participant{}
	if ch.DB.Where("room_id = ?", roomID).First(&room).RecordNotFound() {
		log.Error("NOT FOUND")
		c.AbortWithStatusJSON(404, []string{})
		return
	} else {
		if ch.DB.Model(&room).Association("Participants").Find(&participants).Error != nil {
			log.Error("NOT FOUND")
			c.AbortWithStatusJSON(404, []string{})
			return
		} else {
			room.Participants = participants
		}
	}
	c.JSON(200, &room)
}
