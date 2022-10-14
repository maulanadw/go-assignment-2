package model

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string `gorm:"column:customer_name;type:varchar;size:255"`
	Items        []Item
	OrderedAt    time.Time `gorm:"column:ordered_at;type:timestamp"`
}
