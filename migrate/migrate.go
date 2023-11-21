package main

import (
	"fmt"
	"go-react-example-api/db"
	"go-react-example-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
