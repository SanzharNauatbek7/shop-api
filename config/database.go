package config

//
//import (
//	"log"
//	"shop-api/models"
//
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//)
//
//var DB *gorm.DB
//
//func InitDB() {
//	dsn := "host=localhost port=5432 dbname=shop user=postgres password=true sslmode=disable"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to connect to database:", err)
//	}
//
//	db.AutoMigrate(&models.Product{}) // Автоматическое создание таблицы
//	DB = db
//}
