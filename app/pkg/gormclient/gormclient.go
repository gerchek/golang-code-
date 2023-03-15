package gormclient

import (
	"fmt"
	"log"

	models "project/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type gormConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// NewPgConfig creates new pg config instance
func NewGormConfig(username string, password string, host string, port string, database string) *gormConfig {
	return &gormConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

// NewClient
func NewClient(cfg *gormConfig) (client *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Xml{})

	return db, err
}

//CloseDatabaseConnection method is closing a connection between app and database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
