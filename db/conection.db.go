package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/mrtrom/go-graphql-example-api/config"
	"go.uber.org/zap"
)

// CreateConnetion creates a mysql connection
func CreateConnetion(config *config.Config, log *zap.SugaredLogger) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Error("MySQL is not connected! :(")
		log.Panic(err.Error())
	}

	return db
}
