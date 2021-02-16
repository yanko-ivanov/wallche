package config

import (
	"gorm.io/gorm"
)

// Config type for the application
type Config struct {
	db gorm.Db
}
