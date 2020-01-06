package tools

import (
	"fmt"

	_ "github.com/blog/library/db"
	"github.com/blog/models"
	"github.com/urfave/cli"
)

var InitDB *cli.Command

func init() {
	InitDB = &cli.Command{
		Name:   "migrate",
		Usage:  "Creat tables in database",
		Action: createTables,
	}
}

func createTables(c *cli.Context) error {
	fmt.Println("create tables...")
	models.CreateTable()
	return nil
}
