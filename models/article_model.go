package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	gorm.Model
	ClickQuantity  int
	Content        string
	CreationDate   time.Time
	Introduction   string
	ModifiedDate   time.Time
	Title          string
	Version        int64
	User           User           `gorm:"ForeignKey:ID"`
	Classification Classification `gorm:"ForeignKey:ID"`
	Labels         []Label        `gorm:"many2many:label_article;"`
}

func GetArticle(id int64) (article Article) {
	find := Db.Find(&article, id)
	bytes, _ := json.Marshal(find)
	fmt.Printf(string(bytes))
	a := Article{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), DeletedAt: &time.Time{}}}
	a.ID = 1
	Db.Model(&Article{}).Update("ID", gorm.Expr("11"))
	return
}
