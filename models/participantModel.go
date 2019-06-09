package models

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Participant struct {
	RoomRefer uint `json:"-"`
	//RecordHash string `json:"-" gorm:"unique;type:varchar(255);not null;primary_key; default: null"`
	RoomID    string `json:"room_id" gorm:"type:varchar(255);not null;"`
	UserID    string `json:"user_id" gorm:"not null;default: null"`
	IsBlocked bool   `json:"is_blocked" gorm:"not null;default:'false'"`
	BaseModel
}

func (p *Participant) BeforeSave(scope *gorm.Scope) error {
	log.Info("SAVING PARTICIPANTSSSS: ", p)
	return nil
}
