package model

import (
	"github.com/jinzhu/gorm"
)

type Invoice struct {
	gorm.Model
	Name   string   `gorm:"type:varchar(128);not null" json:"name" form:"name" binding:"required"`
	Client string   `gorm:"type:varchar(128)" json:"client" form:"client"`
	Detail []Detail `json:"detail" form:"detail"`
	Total  int
}

type Detail struct {
	InvoiceID int  `sql:"type:int" gorm:"primary_key"`
	Item      Item `json:"item" form:"item"`
	ItemID    int  `json:"item_id" form:"item_id" sql:"type:int" gorm:"primary_key"`
	Count     int  `json:"count" form:"count" binding:"required,min=1"`
	SubTotal  int
}

type Item struct {
	ID          int    `gorm:"primary_key" json:"id" form:"id"`
	Name        string `gorm:"type:varchar(128);primary_key" json:"name" form:"name" binding:"required"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Unit        string `gorm:"type:varchar(128)" json:"unit" form:"unit" binding:"required"`
	Description string `gorm:"type:varchar(512)" json:"description" form:"description"`
}
