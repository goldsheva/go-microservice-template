package database

import (
	"sync"

	"github.com/goldsheva/go-microservice-template/internal/configs"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	initOnce  sync.Once
	initMutex sync.Mutex
)

func doInit() *gorm.DB {
	config := configs.GetEnvConfig()

	var db *gorm.DB
	switch config.DB_DRIVER {
	case "sqlite3":
		db = InitSQLite3()
	case "mysql":
		db = InitMySQL()
	case "postgres":
		db = InitPostgreSQL()
	default:
		db = InitMySQL()
	}

	logrus.WithFields(logrus.Fields{"gopher": "main", "driver": config.DB_DRIVER}).Info("Database connection established")
	return db
}

func InitDB() *gorm.DB {
	initOnce.Do(func() {
		initMutex.Lock()
		DB = doInit()
		initMutex.Unlock()
	})
	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		InitDB()
	}

	if DB == nil {
		return nil
	}

	sqlDB, err := DB.DB()
	if err != nil {
		initMutex.Lock()
		DB = doInit()
		initMutex.Unlock()
		return DB
	}

	if err := sqlDB.Ping(); err != nil {
		initMutex.Lock()
		DB = doInit()
		initMutex.Unlock()
	}

	return DB
}
