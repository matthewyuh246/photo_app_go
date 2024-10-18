package main

import (
	"fmt"

	"github.com/matthewyu246/back/db"
	"github.com/matthewyu246/back/models"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.Photo{})
}
