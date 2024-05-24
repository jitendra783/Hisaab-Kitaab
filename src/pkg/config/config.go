package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

const (
	USERID    = "userId"
	REQUESTID = "requestID"
	// UCC         = "ucc"
	TOKEN   = "token"
	XLENGTH = "X-length"
	// SCOPE       = "scope"
	 AccessToken = "accessToken"
	// IV256       = "iv"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func LoadConfig(env string, configPaths ...string) {
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	config.AddConfigPath("../app/")
	config.AddConfigPath(".")
	if len(configPaths) != 0 {
		for _, path := range configPaths {
			config.AddConfigPath(path)
		}
	}
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file ", err)
	}
	if env == "server" {
		log.Println("server running in prod getting values from env")
		for _, configKey := range config.AllKeys() {
			if envVal, ok := os.LookupEnv(config.GetString(configKey)); ok {
				log.Println("updating config value with env value for key=", configKey, " value=", envVal)
				config.Set(configKey, envVal)
			} else {
				log.Println("config value not found in env. key= ", configKey)
				os.Exit(1)
			}
		}
	}
	if err == nil {
		log.Println("successfully read server config. values are :", config.AllSettings())
	}

}

func GetConfig() *viper.Viper {
	return config
}
