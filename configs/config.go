package configs

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Config struct {
	DockerAPIGroup string `mapstructure:"DOCKER_API_GROUP"`
	AppPort        string `mapstructure:"APP_PORT"`
}

func LoadConfig(echoLogger echo.Logger, path string) (config Config, err error) {
	// Set the path and name for the .env file
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	// Automatically load environment variables
	viper.AutomaticEnv()

	// Attempt to read the .env file
	err = viper.ReadInConfig()
	// If the .env file doesn't exist or can't be read, log the issue but continue
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// .env file not found; fall back to environment variables
			echoLogger.Info("No .env file found, loading environment variables from host instead")

		} else {
			// For any other error, log and return it
			echoLogger.Error("Error reading .env file: ", err)
			return config, err
		}
	} else {
		// Log that the .env file was loaded successfully
		echoLogger.Info("Successfully loaded .env file")
	}

	err = viper.BindEnv("APP_PORT")
	if err != nil {
		echoLogger.Error("Cannot bind `APP_PORT` environment variable")
	}
	if viper.GetString("APP_PORT") == "" {
		echoLogger.Error("Missing `APP_PORT` environment variable")
	}

	err = viper.BindEnv("DOCKER_API_GROUP")
	if err != nil {
		echoLogger.Error("Cannot bind `DOCKER_API_GROUP` environment variable")
	}
	if viper.GetString("DOCKER_API_GROUP") == "" {
		echoLogger.Info("No specific `DOCKER_API_GROUP` will set it to `default`")
		config.DockerAPIGroup = "default"
	}

	// Unmarshal the loaded configuration into the Config struct
	err = viper.Unmarshal(&config)
	if err != nil {
		echoLogger.Error("Error unmarshaling config: ", err)
		return config, err
	}

	// Log the loaded configuration
	echoLogger.Infof("Loaded configuration: %+v", config)

	return config, nil
}
