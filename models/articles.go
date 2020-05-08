package models

import (
	"fmt"
	"time"

	"github.com/kxiaong/blog/library/db"
)

type Article struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"-"`
	Title     string    `gorm:"type:varchar(256);unique_index" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Status    bool      `gorm:"not null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

func (a *Article) Save() error {
	return db.DB.Save(a).Error
}

func CreateArticle(a Article) error {
	fmt.Println(a.Title)
	fmt.Println(a.Content)

	return db.DB.Create(&Article{
		ID:        1,
		Title:     a.Title,
		Content:   a.Content,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}).Error
}

func GetArticleList() (*[]Article, error) {
	var articles []Article
	err := db.DB.Find(&articles).Error
	return &articles, err
}
