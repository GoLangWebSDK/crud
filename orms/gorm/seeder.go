package gorm

import (
	"fmt"

	"github.com/GoLangWebSDK/records/database"
	"gorm.io/gorm"
)

type Seeder struct {
	DB *database.Database
	Seeders []database.ModelSeeder
	grom *gorm.DB
}

var _ database.ORMSeeder = Seeder{}

func NewSeeder(db *database.Database) (*Seeder, error) {
	seeder := &Seeder{DB: db}
	
	orm, err := db.Adapter.Gorm()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	seeder.grom = orm
	return seeder, nil
}

func (s Seeder) AddSeeder(seeders ...database.ModelSeeder) database.ORMSeeder {
	s.Seeders = append(s.Seeders, seeders...)
	return s
}

func (s Seeder) Run() error {
	for _, seeder := range s.Seeders {
		err := seeder.SeedModel(s.DB)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}