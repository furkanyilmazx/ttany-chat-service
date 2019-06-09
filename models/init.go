package models

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
	if db.DropTable(&Participant{}, &Room{}).Error != nil {
		log.Error("Patlaidkkkk")
	}
	db.AutoMigrate(&Participant{}, &Room{})

	room := Room{
		RoomID:  "2671c20b-0b09-4648-8f4c-0369b284e9b4",
		AdminID: "8fcc9a26-04d0-4f40-8eaf-3d705669acf6",
		Name:    "Sohbet muhabbet",
		Type:    "direct",
		Status:  "active",
		Participants: []Participant{
			{
				RoomID: "2671c20b-0b09-4648-8f4c-0369b284e9b4",
				UserID: "8fcc9a26-04d0-4f40-8eaf-3d705669acf6",
			},
			{
				RoomID: "2671c20b-0b09-4648-8f4c-0369b284e9b4",
				UserID: "cb235d43-e300-4ba6-99be-390ce0812a85",
			},
			{
				RoomID: "2671c20b-0b09-4648-8f4c-0369b284e9b4",
				UserID: "cb235d43-e300-4ba6-99be-390ce0812a85",
			},
		},
	}

	if result := db.Save(&room); result.Error != nil {
		log.Error("Error occured", result.Error)
	}
	return db
}
