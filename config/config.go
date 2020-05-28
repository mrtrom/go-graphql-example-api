package config

import (
	"log"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

type viperLog struct {
	Level            string
	OutputPaths      []string
	ErrorOutputPaths []string
}

type Config struct {
	AppName     string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	JWTSecret   string
	JWTExpireIn time.Duration
	DebugMode   bool
	Log         *viperLog
}

// LoadConfig loads main app configuration
func LoadConfig(path string) *Config {
	config := viper.New()
	config.SetConfigName("development")
	config.AddConfigPath(path)
	err := config.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Could not find any config (.env) file")
		} else {
			log.Fatalf("Fatal error config context file: %s \n", err)
		}
	}

	debugBool, _ := strconv.ParseBool(config.GetString("LOG_DEBUG_MODE"))

	logConfig := &viperLog{
		Level:            config.GetString("LOG_LEVEL"),
		OutputPaths:      []string{"stdout", config.GetString("LOG_OUTPUT_PATH")},
		ErrorOutputPaths: []string{"stderr", config.GetString("LOG_ERROR_PATH")},
	}

	return &Config{
		AppName: config.GetString("APP_NAME"),

		DBHost:     config.GetString("DB_HOST"),
		DBPort:     config.GetString("DB_PORT"),
		DBUser:     config.GetString("DB_USER"),
		DBPassword: config.GetString("DB_PASSWORD"),
		DBName:     config.GetString("DB_NAME"),

		JWTSecret:   config.GetString("AUTH_JWT_SECRET"),
		JWTExpireIn: config.GetDuration("AUTH_JWT_EXPIRE_IN"),

		DebugMode: debugBool,
		Log:       logConfig,
	}
}
