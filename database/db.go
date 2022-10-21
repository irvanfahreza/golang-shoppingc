package database

import (
	"fmt"

	//"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

// github.com/mattn/go-sqlite3
//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

var Db *gorm.DB
func InitDb() *gorm.DB { // OOP constructor
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err !=nil {
		fmt.Println("Error...")
		return nil
	}
	return db
}
