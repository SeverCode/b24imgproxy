package AppConfig

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// setup enviroment
	viper.SetConfigFile("config.yaml")
	viper.AutomaticEnv()
	log.Info("Accessing config " + viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func Get(param string) string {
	return viper.GetString("app." + param)
}