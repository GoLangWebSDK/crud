package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type DBConfig struct {
	DSN    string
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort int
}

type Adapter interface {
	Gorm() gorm.Dialector
}

type Migrator interface {
	Models() []interface{}
	Gorm() []*gormigrate.Migration
}



