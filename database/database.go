package database

import (
	"fmt"
	"log"
	"os"

	"github.com/amusfq/go-fiber/config"
	"github.com/amusfq/go-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	var err error

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_HOST"),
		config.Config("DB_PORT"),
		config.Config("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(1)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.Todo{})
	log.Println("Database Migrated")

	DB = Dbinstance{
		Db: db,
	}
}
