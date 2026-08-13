package database

import (
	"fmt"
	"log"

	"github.com/aditya-singh-finbox/todo-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	fmt.Printf("Host: %s\n", cfg.DBHost)
	fmt.Printf("Port: %s\n", cfg.DBPort)
	fmt.Printf("User: %s\n", cfg.DBUser)
	fmt.Printf("DB: %s\n", cfg.DBName)
	fmt.Printf("Password Length: %d\n", len(cfg.DBPassword))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Ping()
	if err != nil {
		return err
	}
	DB = db
	log.Println("Connected to the database successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
