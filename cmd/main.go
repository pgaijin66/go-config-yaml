package main

import (
	"fmt"
	"log"

	"github.com/pgaijin66/go-config-yaml/config"
)

const (
	configPath = "/Users/pgaijin66/youtube-demo/go/config-demo/data"
	configName = "config"
	ConfigType = "yaml"
)

func main() {
	config, err := config.LoadConfiguration(configPath, configName, ConfigType)
	if err != nil {
		log.Fatalf("could not load configuration file: %v", err)
	}
	fmt.Println(config.Server.Host)
}
