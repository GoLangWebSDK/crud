package gorm

import (
	"fmt"

	"github.com/GoLangWebSDK/records/database"
	"gorm.io/gorm"
)

type GormSeeder struct {
	db *database.Database
	gorm *gorm.DB
	seeders []database.ModelSeeder
}

func NewGormSeeder(db *database.Database) *GormSeeder {
	var err error
	seeder := &GormSeeder{
		db: db,
	}

	seeder.gorm, err = gorm.Open(db.Adapter.Gorm(), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return seeder
}

func (s *GormSeeder) AddSeeder(seeders ...database.ModelSeeder) *GormSeeder {
	s.seeders = append(s.seeders, seeders...)
	return s
}

func (s *GormSeeder) Run() error {
	for _, seeder := range s.seeders {
		err := seeder.SeedModel(s.db)
		if err != nil {
			// tmp error handling...
			panic(err)
		}
	}
	return nil
}