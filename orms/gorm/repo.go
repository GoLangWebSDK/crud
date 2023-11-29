package gorm

import (
	"fmt"

	"github.com/GoLangWebSDK/records"
	"github.com/GoLangWebSDK/records/database"
	"gorm.io/gorm"
)


var _ records.Repository[any] = (*Repository[any])(nil)

type Repository[T any] struct {
	gorm 					 *gorm.DB
	Model 				 T
	deletedAtQuery string
}

func NewRepository[T any](db *database.Database, model T) *Repository[T] {
	orm, err := db.Adapter.Gorm()
	if err != nil {
		fmt.Println(err)
	}

	return &Repository[T]{
		gorm: orm,
		Model: model,
		deletedAtQuery: "%s.deleted_at IS NULL",
	}
}

func (repo *Repository[T]) All() ([]T, error) {
	results := []T{}
	err := repo.gorm.Find(&results).Error
	if err != nil {
		return nil, err
	}
	
	return results, nil
}

func (repo *Repository[T]) Create(model T) error {
	return repo.gorm.Create(&model).Error
}

func (repo *Repository[T]) Read(ID uint32) (T, error) {
	err := repo.gorm.Where("id = ?", ID).First(&repo.Model).Error
	if err != nil {
		return repo.Model, err
	}
	
	return repo.Model, nil	
}

func (repo *Repository[T]) Update(ID uint32, model T) error {
	if ID == 0 {
		return fmt.Errorf("Missing ID")
	}

	err := repo.gorm.First(&repo.Model, ID).Error
	if err != nil {
		return err
	}

	return repo.gorm.Model(&model).Where("id = ?", ID).Updates(model).Error
}

func (repo *Repository[T]) Delete(ID uint32) error {
	if ID == 0 {
		return fmt.Errorf("Missing ID")
	}
	
	err := repo.gorm.First(&repo.Model, ID).Error
	if err != nil {
		return err
	}

	return repo.gorm.Delete(&repo.Model, ID).Error
}
