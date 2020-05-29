package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/mrtrom/go-graphql-example-api/config"
	"github.com/mrtrom/go-graphql-example-api/model"
	"go.uber.org/zap"
)

// CreateConnetion creates a mysql connection
func CreateConnetion(config *config.Config, log *zap.SugaredLogger) *gorm.DB {
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		log.Error("MySQL is not connected! :(")
		log.Panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&model.User{})

	return db
}
