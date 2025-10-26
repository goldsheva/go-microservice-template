package database

import (
	"os"
	"path/filepath"

	"github.com/goldsheva/go-microservice-template/internal/configs"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite3() *gorm.DB {
	config := configs.GetEnvConfig()

	dbPath := config.DB_NAME
	if !filepath.IsAbs(dbPath) {
		wd, err := os.Getwd()
		if err != nil {
			logrus.Panic(err)
		}
		dbPath = filepath.Join(wd, dbPath)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger:                                   gorm_logrus.New(),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatal("Can't init SQLite connection: ", err)
	}

	return db
}
