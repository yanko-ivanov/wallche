package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDb - Initiates a database connection
func InitDb() *gorm.DB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	// "wallche:wallchepass@tcp(db:3306)/wallche?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPass, dbHost, dbPort, dbDatabase)

	// println(dsn)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}
	defer sqlDB.Close()

	return conn
}
