package model

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"column:item_code;type:varchar;size:255;unique"`
	Description string `gorm:"column:description;type:varchar;size:255"`
	Quantitiy   int    `gorm:"column:quantitiy;type:int"`
	OrderID     int    `gorm:"column:order_id;type:int"`
}
