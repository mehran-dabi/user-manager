package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type Configs struct {
	Service  ServiceConfigs
	Database DatabaseConfigs
}

type ServiceConfigs struct {
	Port string
}

type DatabaseConfigs struct {
	Name         string `mapstructure:"name"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	Pass         string `mapstructure:"pass"`
	Driver       string `mapstructure:"driver"`
	RootUser     string `mapstructure:"root_user"`
	RootPassword string `mapstructure:"root_pass"`
}

func Init() *Configs {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)

	viper.SetConfigName("config")
	viper.AddConfigPath(basePath)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	var configs Configs
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading configs: %s", err)
	}

	err := viper.Unmarshal(&configs)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &configs
}
