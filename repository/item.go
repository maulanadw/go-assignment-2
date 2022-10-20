package repository

import (
	"go-assignment-2/model"
	"go-assignment-2/param"

	"gorm.io/gorm"
)

type ItemRepository interface {
	CreateItem(item []model.Item) error
	GetItem() ([]model.Item, error)
	UpdateItemByOrderID(orderId int, item []model.Item) error
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
	var item param.Item
	err := i.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("order_id=?", orderID).Delete(&item).Error
		if err != nil {
			return err
		}

		err = tx.Create(&items).Error
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (i *itemRepository) DeleteItem(ID int) error {
	var item model.Item
	return i.DB.Where("id=?", ID).Delete(&item).Error
}
