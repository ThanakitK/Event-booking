package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App AppConfig `mapstructure:"app"`
	Jwt JwtConfig `mapstructure:"jwt"`
	DB  DBConfig  `mapstructure:"db"`
}

type JwtConfig struct {
	AccessSecret  string `mapstructure:"access_secret"`
	RefreshSecret string `mapstructure:"refresh_secret"`
}
type AppConfig struct {
	Env  string `mapstructure:"env"`
	Port string `mapstructure:"port"`
	Cors string `mapstructure:"cors"`
}

type DBConfig struct {
	URI  string `mapstructure:"uri"`
	Name string `mapstructure:"name"`
}

var Env Config

func NewAppInitEnvironment() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	if err := viper.Unmarshal(&Env); err != nil {
		log.Fatalf("Unable to unmarshal config: %s", err)
	}

	log.Println("config.yaml loaded successfully.")
}
