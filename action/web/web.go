package web

import (
	"github.com/blog/conf"
	"github.com/blog/controllers"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var Web *cli.Command

func init() {
	Web = &cli.Command{
		Name:  "web",
		Usage: "blog website server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "conf, c",
				Value: "",
				Usage: "config file",
			},
			&cli.StringFlag{
				Name:  "args, a",
				Value: "",
				Usage: "multiconfig cmd line args",
			},
		},
		Action: runWeb,
	}
}

func runWeb(c *cli.Context) error {
	LetsRocknRoll()
	return nil
}

func LetsRocknRoll() {
	engine := gin.Default()
	engine.Static("/static", conf.C.Web.WebStaticDir)
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", controllers.Index)

	ArticleApi := engine.Group("/article")
	{
		ArticleApi.GET("/list", controllers.ArticleList)
		ArticleApi.POST("/create", controllers.CreateArticle)
	}

	engine.Run(conf.C.Web.ListenPort)
}
