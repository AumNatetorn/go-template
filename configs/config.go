package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var conf *Config

func Init(path string) error {
	if path == "" {
		path = "configs"
	}

	initViper(path)
	if err := loadConfigs(); err != nil {
		return fmt.Errorf("load config: %v", err)
	}

	if err := conf.Validate(); err != nil {
		return fmt.Errorf("validate config: %v", err)
	}

	return nil
}

func GetConfig() *Config {
	if conf == nil {
		if err := loadConfigs(); err != nil {
			panic(err)
		}
	}
	return conf
}

func initViper(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("cannot read config file: %s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}

func loadConfigs() error {
	_ = viper.Unmarshal(&conf)

	return nil
}
