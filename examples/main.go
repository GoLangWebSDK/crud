package main

import (
	"fmt"

	"github.com/GoLangWebSDK/records/database"
	"github.com/GoLangWebSDK/records/database/adapters"
	"github.com/GoLangWebSDK/records/orms/gorm"
)

type User struct {
	ID        uint32
	FirstName string
	LastName  string
}

func main() {

	adapter := adapters.NewSQLite(database.WithDSN("sqlite.db"))
	db := database.New(adapter)

	users := gorm.NewRepository[User](db, User{})

	allUsers, err := users.All()

	if err != nil {
		fmt.Println(err)
	}

	for _, user := range allUsers {
		fmt.Println(user)
	}
}