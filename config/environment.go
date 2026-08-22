package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App        AppConfig        `mapstructure:"app"`
	Jwt        JwtConfig        `mapstructure:"jwt"`
	DB         DBConfig         `mapstructure:"db"`
	R2         R2Config         `mapstructure:"r2"`
	Push       PushConfig       `mapstructure:"push"`
	Firebase   FirebaseConfig   `mapstructure:"firebase"`
	CloudTasks CloudTasksConfig `mapstructure:"cloud_tasks"`
	MasterDB   MasterDBConfig  `mapstructure:"master_db"`
}

type CloudTasksConfig struct {
	ProjectID       string `mapstructure:"project_id"`
	LocationID      string `mapstructure:"location_id"`
	Queue           string `mapstructure:"queue"`
	InternalBaseURL string `mapstructure:"internal_base_url"`
	InternalSecret  string `mapstructure:"internal_secret"`
}

type FirebaseConfig struct {
	ServiceAccountPath string `mapstructure:"service_account_path"`
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
	URI      string `mapstructure:"uri"`
	Name     string `mapstructure:"name"`
	TenantID string `mapstructure:"tenant_id"` // canonical tenant ID for this deployment
}

type MasterDBConfig struct {
	Name string `mapstructure:"name"` // DB name on the shared cluster
}

type R2Config struct {
	PublicUrl   string `mapstructure:"public_url"`
	AccessKey   string `mapstructure:"access_key"`
	SecretToken string `mapstructure:"secret_token"`
	Endpoint    string `mapstructure:"endpoint"`
	Bucket      string `mapstructure:"bucket"`
}

type PushConfig struct {
	VapidPublicKey  string `mapstructure:"vapid_public_key"`
	VapidPrivateKey string `mapstructure:"vapid_private_key"`
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
