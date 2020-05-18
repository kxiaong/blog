package web

import (
	"github.com/kxiaong/blog/conf"
	"github.com/kxiaong/blog/controllers"
	"github.com/kxiaong/blog/library/db"
	"github.com/kxiaong/blog/library/log"
	"github.com/kxiaong/blog/library/redis"

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
	db.Init()
	redis.Init()
	log.Init()

	// All check, Let's rock and roll!
	LetsRocknRoll()
	return nil
}

func LetsRocknRoll() {
	engine := gin.Default()
	engine.Static("/static", conf.C.Web.WebStaticDir)
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", controllers.Index)
	ArticleApi := engine.RouterGroup.Group("/article")
	{
		ArticleApi.GET("/list", controllers.ArticleList)
		ArticleApi.POST("/create", controllers.CreateArticle)
	}

	PostApi := engine.RouterGroup.Group("/post")
	{
		PostApi.GET("/list", controllers.ListPost)
		PostApi.POST("/create", controllers.CreatePost)
	}

	engine.Run(conf.C.Web.ListenPort)
}
