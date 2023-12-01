package adapters

import (
	"fmt"

	"github.com/GoLangWebSDK/records/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySql struct {
	config *database.DBConfig
}

func NewMySQL(options ...database.DatabaseOptions) *MySql {
	adapter := &MySql{
		config: &database.DBConfig{},
	}

	for _, option := range options {
		option(adapter.config)
	}

	if adapter.config.DSN == "" && adapter.config.DBName == "" {
		fmt.Println("Missing DSN or database configuration for MySQL adapter.") 
		return nil
	}

	return adapter
}

func (adapter *MySql) Gorm() gorm.Dialector {
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	if adapter.config.DSN == "" {
		adapter.config.DSN = fmt.Sprintf(dsn,
			adapter.config.DBUser,
			adapter.config.DBPass,
			adapter.config.DBHost,
			adapter.config.DBPort,
			adapter.config.DBName,
		)
	}	
	return mysql.Open(dsn)
}