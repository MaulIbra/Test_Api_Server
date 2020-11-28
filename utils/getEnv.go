package utils

import (
	"github.com/spf13/viper"
	"log"
)

func GetEnv(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", key, defVal)
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion %v",key)
	}
	return value
}
