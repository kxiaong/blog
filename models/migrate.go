package models

import (
	"fmt"
	"github.com/blog/library/db"
)

func CreateTable() {
	fmt.Println("migrating...")
	db.DB.AutoMigrate(&Article{})
}
