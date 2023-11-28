package config

import "os"

// AppConfig holds the configuration values for the application
type AppConfig struct {
	MongoURI string
	Port     string
}

// LoadConfig loads the application configurations
func LoadConfig() AppConfig {
	return AppConfig{
		MongoURI: os.Getenv("MONGO_URI"),
		Port:     os.Getenv("PORT"),
	}
}
