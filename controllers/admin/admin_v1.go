package admin

import (
	"fmt"

	"mysite/library/log"
	"mysite/models"

	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var (
	logger     = logging.MustGetLogger("controller/admin")
	timeFormat = "2020-01-05 15:06:30"
)

func GetArticlesList(c *gin.Context) {
	requestId := c.MustGet("requestId").(log.RequestID)
	fmt.Println(requestId)
	models.GetArticleList()

}
