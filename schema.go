package records

import (
	"github.com/GoLangWebSDK/records/database"
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

type DBAdapter interface {
	Gorm() (gorm.Dialector, error)
}

type DBMigrator interface {
	Run() error
	AddMigrations(DBMigrations)
	AddModels([]interface{}) 
}

type DBMigrations interface {
	Models() []interface{}
	GormMigrations() []*gormigrate.Migration
}

type ModelSeeder interface {
	SeedModel(*database.Database) error
}

type ORMSeeder interface {
	Run() error
	AddSeeder(seeders ...ModelSeeder) ORMSeeder
}

type Repository[T any] interface {
	All() ([]T, error)
	Create(model T) error
	Read(ID uint32) (T, error)
	Update(ID uint32, model T) error
	Delete(ID uint32) error
}