package configs

import (
	"os"
	"strconv"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sirupsen/logrus"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	LogLevel    logrus.Level
	HTTP_PORT   int
	DB_DRIVER   string
	DB_HOST     string
	DB_PORT     int
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

func GetEnvConfig() *Config {
	once.Do(func() {
		config = &Config{}

		switch os.Getenv("LOG_LEVEL") {
		case "debug":
			config.LogLevel = logrus.DebugLevel
		case "info":
			config.LogLevel = logrus.InfoLevel
		case "warn":
			config.LogLevel = logrus.WarnLevel
		case "error":
			config.LogLevel = logrus.ErrorLevel
		case "fatal":
			config.LogLevel = logrus.FatalLevel
		case "panic":
			config.LogLevel = logrus.PanicLevel
		default:
			config.LogLevel = logrus.InfoLevel
		}

		httpPort, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
		config.HTTP_PORT = httpPort
		config.DB_DRIVER = os.Getenv("DB_DRIVER")
		config.DB_HOST = os.Getenv("DB_HOST")
		dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
		config.DB_PORT = dbPort
		config.DB_NAME = os.Getenv("DB_NAME")
		config.DB_USER = os.Getenv("DB_USER")
		config.DB_PASSWORD = os.Getenv("DB_PASSWORD")

		if err := validation.ValidateStruct(config,
			validation.Field(&config.HTTP_PORT, validation.Required, validation.Min(1), validation.Max(65535)),
			validation.Field(&config.DB_DRIVER, validation.Required, validation.In("sqlite3", "mysql", "postgres")),
			validation.Field(&config.DB_HOST, validation.Length(1, 100)),
			validation.Field(&config.DB_PORT, validation.Min(1), validation.Max(65535)),
			validation.Field(&config.DB_NAME, validation.Required, validation.Length(4, 64)),
			validation.Field(&config.DB_USER, validation.Length(0, 64)),
			validation.Field(&config.DB_PASSWORD, validation.Length(0, 64)),
		); err != nil {
			logrus.Fatalf("Can't parse config.env: %v", err)
		}
	})

	return config
}
