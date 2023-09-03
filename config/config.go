package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT string

	JWT_SECRET        string
	JWT_EXPIRE        string
	JWT_COOKIE_EXPIRE string

	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_HOST     string
	DB_PORT     string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err)
	}

	log.Println("[APP] Load server successfully")
}
