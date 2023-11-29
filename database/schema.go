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
	GetDSN() string
	Gorm() (*gorm.DB, error)
}

type ORM[T any] interface {
	Init() error
	GetAllRecords() ([]T, error)
}

type ModelSeeder interface {
	GormSeeder(gorm *gorm.DB) error
}

type ORMSeeder interface {
	Run() error
	AddSeeder(seeders ...ModelSeeder) ORMSeeder
}

type Migrator interface {
	Models() []interface{}
	Gorm() []*gormigrate.Migration
}



