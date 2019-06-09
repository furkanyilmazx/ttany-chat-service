package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Rooms []Room

func (rs *Rooms) AllRooms() (err error) {
	return db.Preload("Participants").Find(&rs).Error
}

type Room struct {
	RoomID       string        `json:"room_id" gorm:"unique;not null"`
	AdminID      string        `json:"admin_id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`   // 'direct|group'
	Status       string        `json:"status"` // 'closed|fraud' and so on
	Participants []Participant `json:"participants" gorm:"foreignkey:RoomRefer"`
	BaseModel
}

func (r *Room) BeforeCreate(scope *gorm.Scope) error {
	uud := uuid.New().String()
	scope.SetColumn("RoomID", uud)
	log.Info("CREATING ROOM ID: ", uud)
	return nil
}

func (r *Room) GetByRoomID(roomID string) (err error) {
	return db.Preload("Participants").Where("room_id = ?", roomID).First(&r).Error
}

func (r *Room) CreateRoom() (err error) {
	if db.NewRecord(r) {
		if err = db.Create(&r).Error; err != nil {
			return err
		}
		return
	}
	log.Error("Not createadddddddddddddddd")
	return errors.New("Not created")
}
