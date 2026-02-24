package config

import(
	"github.com/spf13/viper"
	"log"
)

func ViperEnvConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	
	if err != nil {
		log.Println("There is an error with loading .env: ", err)
	}
	log.Println("Viper Env Configured")
}