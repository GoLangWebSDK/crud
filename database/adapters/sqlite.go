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
	adapter := &SQLite{
		config: &database.DBConfig{},
	}

	for _, option := range options {
		option(adapter.config)
	}

	if adapter.config.DSN == "" && adapter.config.DBName == "" {
		fmt.Println("Missing DSN or database configuration for SQLite adapter.") 
		return nil
	}

	return adapter
}

func (adapter *SQLite) Gorm() (gorm.Dialector, error) {
	var dsn string
	if adapter.config.DSN == "" {
		dsn = adapter.config.DBName
	}	
	return sqlite.Open(dsn), nil
}