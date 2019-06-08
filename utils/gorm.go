package utils

import (
	"github.com/jinzhu/gorm"
)

func LoadGormConfig(db *gorm.DB) {
	db.LogMode(true)
	//db.SetLogger(log.StandardLogger())

}
