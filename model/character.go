package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name string
}
