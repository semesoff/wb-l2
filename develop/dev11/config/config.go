package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type ConfigProvider interface {
	GetConfig() *Config
}

type Config struct {
	App      App      `yaml:"app"`
	Database Database `yaml:"database"`
}

type App struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Driver   string `yaml:"driver"`
}

type ConfigManager struct {
	cfg *Config
}

func NewConfigManager() *ConfigManager {
	cfg := &ConfigManager{
		cfg: &Config{},
	}
	cfg.InitConfig()
	return cfg
}

func (cm *ConfigManager) InitConfig() {
	// Read file content
	data, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	// Close the file after reading
	defer func(data *os.File) {
		if err := data.Close(); err != nil {
			log.Fatalln(err)
		}
	}(data)

	// Unmarshall the data from []byte to AppConfig struct
	decoder := yaml.NewDecoder(data)
	if err := decoder.Decode(cm.cfg); err != nil {
		log.Fatalln(err)
	}

	log.Println("Config is initialized.")
}

func (cm *ConfigManager) GetConfig() *Config {
	return cm.cfg
}
