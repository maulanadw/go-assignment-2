package repository

import (
	"go-assignment-2/model"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item []model.Item) error
	GetItem() ([]model.Item, error)
	UpdateItemByOrderID(orderId int, item []model.Item) error
	DeleteItem(ID int) error
}

type itemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	db.AutoMigrate(&model.Item{})
	return &itemRepository{
		DB: db,
	}
}

func (i *itemRepository) CreateItem(item []model.Item) error {
	return i.DB.Create(&item).Error
}

func (i *itemRepository) GetItem() ([]model.Item, error) {
	var items []model.Item
	err := i.DB.Find(&items).Error
	return items, err
}

func (i *itemRepository) UpdateItemByOrderID(orderID int, items []model.Item) error {
	return i.DB.Model(model.Item{}).Where("order_id = ", orderID).Updates(&items).Error
}

func (i *itemRepository) DeleteItem(ID int) error {
	var item model.Item
	return i.DB.Where("id=?", ID).Delete(&item).Error
}
