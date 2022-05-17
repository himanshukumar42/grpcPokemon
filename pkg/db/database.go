package db

import (
	"log"

	"github.com/jinzhu/gorm"
)

func Init() *gorm.DB {
	// connectionString := utils.GetConnectionString()
	db, err := gorm.Open("sqlite3", "testdb")
	if err != nil {
		log.Fatalf("couldn't connect to database: %v\n", err)
	}
	defer db.Close()

	return db
}
