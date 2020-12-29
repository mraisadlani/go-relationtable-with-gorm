package config

import (
	"fmt"
	"os"

	"github.com/vanilla/go-relationtable/api/entities"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.Debug().AutoMigrate(&entities.User{}, &entities.Role{}, &entities.RoleUser{}, &entities.Tipe{}, &entities.Menu{}, &entities.SubMenu{}, &entities.RoleMenu{}, &entities.UserMenu{})

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
