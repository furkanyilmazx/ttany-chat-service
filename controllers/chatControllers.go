package controllers

import (
	"net/http"
	"ttany-chat-service/models"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func GetChatRooms(c *gin.Context) {
	var rooms = models.Rooms{}
	if err := rooms.AllRooms(); err != nil {
		log.Error("Error Occured", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Debug(rooms)
	c.JSON(http.StatusOK, &rooms)
}

func CreateChatRoom(c *gin.Context) {
	var room = models.Room{}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Info(room)
	if err := room.CreateRoom(); err != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "created"})
}

func GetChatRoomWithRoomID(c *gin.Context) {
	roomID := c.Param("room_id")
	var room = models.Room{}
	if err := room.GetByRoomID(roomID); gorm.IsRecordNotFoundError(err) {
		log.Error("NOT FOUND", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else if err != nil {
		log.Error("Error Occured", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, &room)
}
