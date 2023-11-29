package records

import (
	"github.com/GoLangWebSDK/records/database"
)

type DBConfig struct {
	DSN    string
	DBName string
	DBUser string
	DBPass string
	DBHost string
	DBPort int
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