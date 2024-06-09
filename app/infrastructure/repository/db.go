package infrastructure

import (
	"fmt"
	"log/slog"
	"naive-feed-service/app/config"
	"naive-feed-service/app/util"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var db *gorm.DB

func NewDb(dbConfig *config.Db) *gorm.DB {
	var err error

	once.Do(func() {
		db, err = gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Port)), &gorm.Config{})
	})
	if err != nil {
		util.Logger.Error("Failed to connect to database", slog.Any("err", err))
		os.Exit(1)
	}
	return db
}
