package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Article struct {
	Id               int64  `gorm:"primary_key"`
	ClickQuantity    int    `gorm:"type:bigint(20);"`
	Content          string `gorm:"type:varchar(255);"`
	CreationDate     time.Time
	Introduction     string
	ModifiedDate     time.Time
	Title            string
	Version          int64
	AuthorId         int64
	ClassificationId int64
}

func GetArticle(id int64) (article Article) {
	Db.Find(&article, id)
	return
}
