package db

import (
	"fmt"
	"time"

	"github.com/blog/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	fmt.Println("db init...")
	fmt.Println(conf.C.Database.Postgres)
	var dsn string
	fmt.Println(conf.C.Database.Postgres.Host)
	dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		conf.C.Database.Postgres.Host,
		conf.C.Database.Postgres.Port,
		conf.C.Database.Postgres.User,
		conf.C.Database.Postgres.DBName,
		conf.C.Database.Postgres.Password,
	)

	var err error
	fmt.Println("postgres conn sql: ", dsn)
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = DB.DB().Ping()
	if err != nil {
		fmt.Println("DB ping failed: ", err)
		panic(err)
	}

	DB.DB().SetConnMaxLifetime(30 * time.Second)
	DB.DB().SetMaxIdleConns(200)
	DB.DB().SetMaxOpenConns(100)
	DB.BlockGlobalUpdate(true)
}
