package db

import (
	"fmt"
	"isar-golang-training-beginner/config"
	"isar-golang-training-beginner/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	configuration := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", configuration.DB_HOST, configuration.DB_USERNAME, configuration.DB_PASSWORD, configuration.DB_NAME, configuration.DB_PORT)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB Connection Error")
	}
	db.AutoMigrate(&model.PaymentCodes{})
}

func DbManager() *gorm.DB {
	return db
}
