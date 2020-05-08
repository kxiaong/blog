package controllers

import (
	"fmt"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kxiaong/blog/models"
)

func ArticleList(c *gin.Context) {
	r := make(map[string]string)
	r["a"] = "article 1"
	r["b"] = "article 2"
	r["c"] = "article 3"
	c.JSON(http.StatusOK, r)

}

func CreateArticle(c *gin.Context) {
	var article models.Article
	article.ID = 1
	article.Content = "this is my first blog article"
	article.Title = "my first blog article"
	article.Status = true
	//if err := c.ShouldBindJSON(&article); err != nil {
	//c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
	//return
	//}

	//article.ID = 1
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()
	fmt.Println("examing article obj------")
	fmt.Println(article.Title)
	fmt.Println(article.Content)
	fmt.Println(article.ID)
	fmt.Println(article.CreatedAt)
	fmt.Println(article.UpdatedAt)
	fmt.Println("end of examing article obj")
	models.CreateArticle(article)
	//if err := article.Save(); err != nil {
	//c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
	//return
	//}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
