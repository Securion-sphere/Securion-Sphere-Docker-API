package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort   int16  `mapstructure:"APP_PORT"`
	JwtSecret string `mapstructure:"JWT_SECRET"`
}

var loadedConfig *Config

// GetConfig safely returns the loaded configuration
func GetConfig() *Config {
	if loadedConfig == nil {
		log.Warn("Config not loaded. Call LoadConfig() first.")
		return &Config{} // Return an empty config to prevent nil pointer errors
	}
	return loadedConfig
}

// LoadConfig initializes and loads the configuration
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env.local") // Load from .env.local
	viper.SetConfigType("env")        // Define config type

	// Set default values
	viper.SetDefault("APP_PORT", 8080)

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("No .env.local file found, using environment variables")
			viper.AutomaticEnv()
			if err := viper.BindEnv("APP_PORT"); err != nil {
				return nil, err
			}
			if err := viper.BindEnv("JWT_SECRET"); err != nil {
				return nil, err
			}
		} else {
			log.Error("Error reading .env.local file:", err)
			return nil, err
		}
	} else {
		log.Info("Successfully loaded .env.local file")
	}

	// Bind values from Viper to Config struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Error("Failed to unmarshal config:", err)
		return nil, err
	}

	if config.JwtSecret == "" {
		log.Error("Configuration error: JWT_SECRET is not set!")
		return nil, fmt.Errorf("JWT_SECRET is not set")
	}

	// Store the loaded config globally
	loadedConfig = &config

	return loadedConfig, nil
}
