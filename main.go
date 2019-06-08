package main

import (
	"ttany-chat-service/middlewares"
	"ttany-chat-service/models"
	"ttany-chat-service/routes"
	"ttany-chat-service/utils"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Infoln("Server initializing...")
	utils.LoadConfig()

	db, err := gorm.Open(viper.GetString("database.driver"), viper.GetString("database.connection"))
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	utils.LoadLogConfig()
	utils.LoadGormConfig(db)

	if db.DropTable(&models.Participant{}, &models.Room{}).Error != nil {
		log.Error("Patlaidkkkk")
	}
	db.AutoMigrate(&models.Participant{}, &models.Room{})

	r := gin.New()

	room := models.Room{
		RoomID:  "2671c20b-0b09-4648-8f4c-0369b284e9b4",
		AdminID: "8fcc9a26-04d0-4f40-8eaf-3d705669acf6",
		Name:    "Sohbet muhabbet",
		Type:    "direct",
		Status:  "active",
		Participants: []models.Participant{
			{
				RoomID: "2671c20b-0b09-4648-8f4c-0369b284e9b4",
				UserID: "8fcc9a26-04d0-4f40-8eaf-3d705669acf6",
			},
			{
				UserID: "cb235d43-e300-4ba6-99be-390ce0812a85",
			},
		},
	}

	if result := db.Save(&room); result.Error != nil {
		log.Error("Error occured", result.Error)
	}

	r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery())

	api := r.Group("api")
	v1 := api.Group("v1")

	routes.LoadRoutes(v1, db)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()

	defer log.Infoln("Server Closed")

}
