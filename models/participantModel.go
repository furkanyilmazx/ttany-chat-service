package models

type Participant struct {
	RoomID    string `json:"room_id" gorm:"type:varchar(255);not null; default: null"`
	UserID    string `json:"user_id" gorm:"not null;primary_key;default: null"`
	IsBlocked bool   `json:"is_blocked" gorm:"not null;default:'false'"`
	BaseModel
}
