package adapters

import (
	"fmt"

	"github.com/GoLangWebSDK/records/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var _ database.Adapter = (*SQLite)(nil)

type SQLite struct {
	config *database.DBConfig
}

func NewSQLite(options ...database.DatabaseOptions) *SQLite {
	adapter := &SQLite{}

	for _, option := range options {
		option(adapter.config)
	}

	if adapter.config.DSN == "" && adapter.config == nil {
		fmt.Println("Missing DSN or database configuration for SQLite adapter.") 
		return nil
	}

	return adapter
}

func (adapter *SQLite) GetDSN() string {
	if adapter.config.DSN != "" {
		return adapter.config.DSN
	}
	return adapter.config.DBName
}

func (adapter *SQLite) Gorm() (*gorm.DB, error) {
	dsn := adapter.GetDSN()
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}