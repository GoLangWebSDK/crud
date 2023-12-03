package gorm

import (
	"fmt"

	"github.com/GoLangWebSDK/records"
	"github.com/GoLangWebSDK/records/database"
	"gorm.io/gorm"
)

var _ records.Repository[any] = (*Repository[any])(nil)

type Repository[T any] struct {
	DB 						 *gorm.DB
	Model 				 T
	deletedAtQuery string
}

func NewRepository[T any](db *database.Database, model T) *Repository[T] {
	repo := &Repository[T]{
		Model: model,
		deletedAtQuery: "%s.deleted_at IS NULL",
	}
	
	gorm, err := gorm.Open(db.Adapter.Gorm(), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	repo.DB = gorm
	return repo
}

func (repo *Repository[T]) All() ([]T, error) {
	results := []T{}
	err := repo.DB.Find(&results).Error
	if err != nil {
		return nil, err
	}
	
	return results, nil
}

func (repo *Repository[T]) Create(model T) error {
	return repo.DB.Create(&model).Error
}

func (repo *Repository[T]) Read(ID uint32) (T, error) {
	err := repo.DB.Where("id = ?", ID).First(&repo.Model).Error
	if err != nil {
		return repo.Model, err
	}
	
	return repo.Model, nil	
}

func (repo *Repository[T]) Update(ID uint32, model T) error {
	if ID == 0 {
		return fmt.Errorf("Missing ID")
	}

	err := repo.DB.First(&repo.Model, ID).Error
	if err != nil {
		return err
	}

	return repo.DB.Model(&model).Where("id = ?", ID).Updates(model).Error
}

func (repo *Repository[T]) Delete(ID uint32) error {
	if ID == 0 {
		return fmt.Errorf("Missing ID")
	}
	
	err := repo.DB.First(&repo.Model, ID).Error
	if err != nil {
		return err
	}

	return repo.DB.Delete(&repo.Model, ID).Error
}