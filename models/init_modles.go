package models

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/kxiaong/blog/conf"
	"github.com/kxiaong/blog/library/db"
)

func CreateTable(noPrompt bool) error {
	dbName := conf.C.Database.Mysql.DB
	if conf.C.Database.Driver == "postgres" {
		dbName = conf.C.Database.Postgres.DBName
	}

	if !noPrompt {
		fmt.Printf("init db(%s-%s), all exists data will be deletedi!\n", conf.C.Database.Driver, dbName)
		fmt.Print("y/n: ")
		reader := bufio.NewReader(os.Stdin)
		if text, _ := reader.ReadString('\n'); text != "y\n" {
			log.Println("cancel!")
			return nil
		}
	}

	// change table, change here
	db.DB.DropTableIfExists(&Article{})
	db.DB.DropTableIfExists(&User{})

	_db := db.DB.LogMode(true)
	if conf.C.Database.Driver == "mysql" {
		_db = db.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	}

	if err := _db.CreateTable(&Article{}).Error; err != nil {
		return err
	}
	if err := _db.CreateTable(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func MigrateTable() error {
	// true for detailed logs, false for error-only logs
	_db := db.DB.LogMode(true)
	if conf.C.Database.Driver == "mysql" {
		_db = db.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")
	}

	// when adding a model, add AutoMigrate here
	if err := _db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	if err := _db.AutoMigrate(&Article{}).Error; err != nil {
		return err
	}

	return nil
}
