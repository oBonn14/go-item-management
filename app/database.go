package app

import (
	"Go-CRUD/app/domain/customer"
	"Go-CRUD/app/domain/item"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var (
	database = os.Getenv("DB_NAME")
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	DB       *gorm.DB
)

func ConnectToDatabase() *gorm.DB {
	if DB != nil {
		return DB
	}
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Gagal konek. \n", err)
		os.Exit(2)
	}

	err = db.AutoMigrate(&customer.Customer{})
	err = db.AutoMigrate(&item.Item{})
	if err != nil {
		log.Fatal("Gagal Migration. \n", err)
	}

	DB = db
	return DB
}
