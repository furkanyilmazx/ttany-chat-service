package models

import (
	log "github.com/sirupsen/logrus"
)

type Room struct {
	RoomID       string        `json:"room_id" gorm:"unique;not null"`
	AdminID      string        `json:"admin_id"`
	Name         string        `json:"name"`
	Type         string        `json:"type"`   // 'direct|group'
	Status       string        `json:"status"` // 'closed|fraud' and so on
	Participants []Participant `json:"participants" gorm:"foreignkey:id"`
	BaseModel
}

func (r *Room) BeforeCreate() (err error) {
	log.Info(r)
	return
}
