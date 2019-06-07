package main

import (
	"ttany-chat-service/utils"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.Infoln("Server initializing...")
	utils.LoadConfig()
	utils.LoadLogConfig()

	log.Infoln("Ddatabase driver: ", viper.GetString("database.driver"))

	defer log.Infoln("Server Closed")

}
