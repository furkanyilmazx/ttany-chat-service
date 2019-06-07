package main

import (
	"ttany-chat-service/middlewares"
	"ttany-chat-service/routes"
	"ttany-chat-service/utils"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.Infoln("Server initializing...")
	utils.LoadConfig()
	utils.LoadLogConfig()

	log.Infoln("Ddatabase driver: ", viper.GetString("database.driver"))

	r := gin.New()
	r.Use(middlewares.LoggerMiddleware())
	r.Use(gin.Recovery())

	routes.LoadRoutes(r)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()

	defer log.Infoln("Server Closed")

}
