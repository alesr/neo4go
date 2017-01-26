package neo4go

import (
	"log"

	"github.com/spf13/viper"
)

const loadConfigErrMsg = "Error loading config file"

// LoadConfig from Viper
func loadConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("%s: %s", loadConfigErrMsg, err)
	}
}
