package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadLogConfig() {
	deployMode := viper.GetString("DEPLOY_MODE")
	if deployMode == "" || deployMode == "prod" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	log.SetReportCaller(true)

}
