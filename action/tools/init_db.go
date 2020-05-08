package tools

import (
	"github.com/kxiaong/blog/library/db"
	"github.com/kxiaong/blog/models"
	"github.com/urfave/cli"
)

var InitDB *cli.Command = &cli.Command{
	Name:   "init_db",
	Usage:  "init datbase",
	Action: runInitDB,
}

var MigrateDB *cli.Command = &cli.Command{
	Name:   "migrate",
	Usage:  "migrate tables",
	Action: runMigrateDB,
}

func runInitDB(c *cli.Context) error {
	db.Init()
	return models.CreateTable(false)
}

func runMigrateDB(c *cli.Context) error {
	db.Init()
	return models.MigrateTable()
}
