package config

import (
	"fmt"
	"gin/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := "host=localhost user=postgres password=1234 dbname=trial_class_go port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err.Error())
	} else {
		fmt.Println("connected to db")
		DB = db
	}

	db.AutoMigrate(&entity.Order{}, &entity.Product{})
}
