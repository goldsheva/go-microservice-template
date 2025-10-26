package database

import (
	"fmt"

	"github.com/goldsheva/go-microservice-template/internal/configs"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() *gorm.DB {
	config := configs.GetEnvConfig()

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC&interpolateParams=true",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)), &gorm.Config{
		Logger:                                   gorm_logrus.New(),
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatal("Can't init MySQL connection: ", err)
	}

	return db
}
