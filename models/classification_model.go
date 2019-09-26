package models

import "github.com/jinzhu/gorm"

type Classification struct {
	gorm.Model
	Name string
}
