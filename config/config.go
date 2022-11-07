package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Configuration struct {
	Environment string

	Server struct {
		Host string
		Port int
	}

	Db struct {
		Password string
	}
}

func LoadConfiguration(configPath, configName, ConfigType string) (*Configuration, error) {
	var config *Configuration

	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(ConfigType)

	viper.AutomaticEnv()
	// first checks for SERVER.PORT
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	// after above instruction, checks for SERVER_PORT

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("could not read the config file: %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal: %v", err)
	}

	return config, nil
}
