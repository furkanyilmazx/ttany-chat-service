package chat

import (
	"errors"
	"time"
	"ttany-chat-service/utils"
	"ttany-chat-service/utils/paginator"

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
	return db.Find(&self).Error
}

func (self *RoomModels) AllRoomsPaginated(limit, offset string) (count uint, err error) {
	db := utils.GetDB()
	return count, db.Limit(limit).Offset(offset).Find(&self).Count(&count).Error
}

func (self *RoomModels) AllRoomsCursorPaginated(p paginator.Paginator) (count uint, cursors paginator.Cursors, err error) {
	db := utils.GetDB()
	result := p.Paginate(db, self)
	cursors = p.GetNextCursors()
	return count, cursors, result.Count(&count).Error
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

func (self *RoomModel) GetByRoomID(roomID string) (err error) {
	db := utils.GetDB()
	return db.Preload("Participants").Where("room_id = ?", roomID).First(&self).Error
}

func (self *RoomModel) CreateRoom() (err error) {
	db := utils.GetDB()
	if db.NewRecord(&self) {
		if err = db.Create(&self).Error; err != nil {
			return err
		}
	}
	return
	log.Error("Not createadddddddddddddddd")
	return errors.New("Not created")
}

type ParticipantModel struct {
	RoomRefer uint64     `json:"-"`
	RoomID    string     `json:"room_id" gorm:"type:char(36);not null;default: null;primary_key"`
	UserID    string     `json:"user_id" gorm:"type:char(36);not null;default: null;primary_key"`
	IsBlocked bool       `json:"is_blocked" gorm:"not null;default:'false'"`
	Accepted  bool       `json:"accepted" gorm:"not null;default:false"` // invite status 'accepted|not_accepted'
	Leaved    bool       `json:"leaved" gorm:"not null;default:false"`   // invite status 'accepted|not_accepted'
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index" `
}
