package controllers

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	a := make(map[string]string)
	a["title1"] = "title1"
	a["date1"] = "2020-5-18"
	a["abstract1"] = "abstract1"
	a["tags1"] = "分布式"
	a["title2"] = "title2"
	a["date2"] = "2020-05-19"
	a["abstract2"] = "abstract2"
	a["tags2"] = "golang"
	arr := make([]map[string]string, 1)
	arr = append(arr, a)

	c.HTML(200, "index.html", gin.H{"arr": arr})
}
