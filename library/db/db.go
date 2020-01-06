package db

import (
	"fmt"
	"time"

	"github.com/blog/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	var dsn string
	if conf.C.Database.Driver == "mysql" {
		dsn = fmt.Sprintf("%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
			conf.C.Database.Mysql.UserPassword,
			conf.C.Database.Mysql.HostPort,
			conf.C.Database.Mysql.DB,
		)
	} else if conf.C.Database.Driver == "postgres" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
			conf.C.Database.Postgres.Host,
			conf.C.Database.Postgres.Port,
			conf.C.Database.Postgres.User,
			conf.C.Database.Postgres.DBName,
			conf.C.Database.Postgres.Password,
		)
	} else {
		panic(fmt.Sprintf("invalid db driver: %s", conf.C.Database.Driver))
	}

	var err error
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	// Is DB ok?
	if err = DB.DB().Ping(); err != nil {
		panic(fmt.Sprintf("Ping db failed: %s", err.Error()))
	}

	DB.DB().SetConnMaxLifetime(30 * time.Second)
	DB.DB().SetMaxIdleConns(200)
	DB.DB().SetMaxOpenConns(100)
	DB.BlockGlobalUpdate(true)
}
