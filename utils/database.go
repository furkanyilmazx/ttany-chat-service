package utils

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	var err error
	db, err = gorm.Open(viper.GetString("database.driver"), viper.GetString("database.connection"))
	db.LogMode(true)
	//db.SetLogger(log.StandardLogger())
	if err != nil {
		log.Error("failed to dbect database", err)
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}
