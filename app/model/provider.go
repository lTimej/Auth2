package model

import (
	"gorm.io/gorm"
)

type Provider struct {
	gorm.Model
	Code string `json:"code" gorm:"type:varchar(64);unique"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}
