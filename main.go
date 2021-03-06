package main

import (
	"os"

	"github.com/kxiaong/blog/action/tools"
	"github.com/kxiaong/blog/action/web"
	"github.com/urfave/cli"
)

var (
	cmds []*cli.Command
)

func init() {
	cmds = []*cli.Command{
		web.Web,
		tools.InitDB,
		tools.MigrateDB,
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "博客站点"
	app.Usage = "博客站点管理平台"
	app.Commands = cmds

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
