// ToDo: Consider adding context to functions for better cancellation and time out support
package gorm

import (
	"fmt"

	"github.com/GoLangWebSDK/crud"
	"github.com/GoLangWebSDK/crud/database"
	"gorm.io/gorm"
)

var _ crud.Repository[any] = (*Repository[any])(nil)

type Repository[T any] struct {
	DB             *gorm.DB
	deletedAtQuery string
}

func NewRepository[T any](db *database.Database) *Repository[T] {
	repo := &Repository[T]{
		DB:             db.Adapter.Gorm(),
		deletedAtQuery: "%s.deleted_at IS NULL",
	}
	return repo
}

func (repo *Repository[T]) Create(newRecord T) (*T, error) {
	err := repo.DB.Create(&newRecord).Error
	if err != nil {
		return nil, err
	}
	return &newRecord, nil
}

func (repo *Repository[T]) ReadAll() ([]T, error) {
	records := []T{}

	err := repo.DB.Find(&records).Error
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (repo *Repository[T]) Read(ID uint) (*T, error) {
	var record T
	err := repo.DB.First(&record, ID).Error
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (repo *Repository[T]) Update(ID uint, updatedFields T) (*T, error) {
	if ID == 0 {
		return nil, fmt.Errorf("invalid ID provided")
	}

	var record T
	err := repo.DB.First(&record, ID).Error
	if err != nil {
		return nil, err
	}

	err = repo.DB.Model(&record).Where("id = ?", ID).Updates(updatedFields).Error
	if err != nil {
		return nil, err
	}

	return &record, nil
}

func (repo *Repository[T]) Delete(ID uint) error {
	if ID == 0 {
		return fmt.Errorf("invalid ID provided")
	}

	var record T
	err := repo.DB.First(&record, ID).Error
	if err != nil {
		return err
	}

	return repo.DB.Delete(&record).Error
}
