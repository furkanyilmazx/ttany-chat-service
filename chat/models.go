package chat

import (
	"errors"
	"hash/fnv"
	"time"
	"ttany-chat-service/utils"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type BaseModel struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index" `
}

type RoomModels []RoomModel

func (self *RoomModels) AllRooms() (err error) {
	db := utils.GetDB()
	return db.Preload("Participants").Find(&self).Error
}

type RoomModel struct {
	RoomID            string             `json:"room_id" gorm:"unique;not null"`
	AdminID           string             `json:"admin_id"`
	Name              string             `json:"name"`
	Type              string             `json:"type"`   // 'direct|group'
	Status            string             `json:"status"` // 'closed|fraud' and so on
	ParticipantModels []ParticipantModel `json:"participants" gorm:"foreignkey:RoomRefer"`
	BaseModel
}

/*
func (self *RoomModel) BeforeCreate(scope *gorm.Scope) error {
	uud := uuid.New().String()
	scope.SetColumn("RoomID", uud)
	log.Info("CREATING ROOM ID: ", uud)
	return nil
} */

func (self *RoomModel) GetByRoomID(roomID string) (err error) {
	db := utils.GetDB()
	return db.Preload("Participants").Where("room_id = ?", roomID).First(&self).Error
}

func (self *RoomModel) CreateRoom() (err error) {
	db := utils.GetDB()
	if db.NewRecord(self) {
		if err = db.Create(&self).Error; err != nil {
			return err
		}
		return
	}
	log.Error("Not createadddddddddddddddd")
	return errors.New("Not created")
}

type ParticipantModel struct {
	RoomRefer uint64     `json:"-"`
	RoomID    string     `json:"room_id" gorm:"type:varchar(255);not null;default: null;primary_key"`
	UserID    string     `json:"user_id" gorm:"not null;default: null;primary_key"`
	IsBlocked bool       `json:"is_blocked" gorm:"not null;default:'false'"`
	Accepted  bool       `json:"accepted" gorm:"not null;default:false"` // invite status 'accepted|not_accepted'
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index" `
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func (p *ParticipantModel) BeforeCreate(scope *gorm.Scope) error {
	log.Info("SAVING PARTICIPANTSSSS: ", p)
	return nil
}
