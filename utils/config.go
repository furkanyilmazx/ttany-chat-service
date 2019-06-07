package utils

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName("env-dev") // name of config file (without extension)
	viper.AddConfigPath(".")       // optionally look for config in the working directory
	viper.AddConfigPath("configs") // path to look for the config file in
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		switch err.(type) {
		default:
			panic(fmt.Errorf("Fatal error loading config file: %s \n", err))
		case viper.ConfigFileNotFoundError:
			log.Errorln("No config file found. Using defaults and environment variables")
		}
	}
	viper.SetEnvPrefix(viper.GetString("env-prefix"))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	for _, key := range viper.AllKeys() {
		log.Infoln("Config File Key:", key, " Value:", viper.Get(key))
	}
	log.Infoln("Key: DEPLOY_MODE Value: ", viper.GetString("DEPLOY_MODE"))

	return nil
}
