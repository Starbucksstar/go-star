package initialize

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"star/src/config"
	"star/src/global"
)

func InitConfig() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	var configuration config.MysqlConfiguration
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}
	global.MysqlConfig = &configuration
}
