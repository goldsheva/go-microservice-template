package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"fmt"

	"github.com/goldsheva/go-microservice-template/internal/configs"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
)

func InitPostgreSQL() *gorm.DB {
	config := configs.GetEnvConfig()

	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		config.DB_HOST,
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_NAME,
		config.DB_PORT,
	)), &gorm.Config{
		Logger:                                   gorm_logrus.New(),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logrus.Fatal("Can't init PostgreSQL connection: ", err)
	}

	return db
}
