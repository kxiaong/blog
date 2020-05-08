package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kxiaong/blog/library/db"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	UID       uuid.UUID `gorm:"type:uuid;primary_key;" json:"uid"`
	Nickname  string    `gorm:"type:varchar(128);" json:"name"`
	Email     string    `gorm:"varchar(256);" json:"email"`
	Authority int       `json:"auth"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	uid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	return scope.SetColumn("UID", uid)
}

func (u *User) Save() error {
	return db.DB.Save(u).Error
}
