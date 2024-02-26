package database

import (
	"log"
	"quiz/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	dsn := "root:@tcp(127.0.0.1:3306)/quiz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// Auto Migration

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Quiz{})
	db.AutoMigrate(&models.Question{})
	db.AutoMigrate(&models.StudentAnswer{})

	return db, nil

}
