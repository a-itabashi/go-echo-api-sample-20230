package main

import (
	"fmt"
	"go-echo-api-sample-202306/db"
	"go-echo-api-sample-202306/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.Task{})
}
