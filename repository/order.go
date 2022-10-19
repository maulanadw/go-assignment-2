package repository

import (
	"go-assignment-2/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrder() ([]model.Order, error)
	UpdateOrder(ID int, order model.Order) error
	DeleteORder(ID int) error
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	db.AutoMigrate(&model.Order{})
	return &orderRepository{
		DB: db,
	}
}

func (or *orderRepository) CreateOrder(order model.Order) (model.Order, error) {
	err := or.DB.Create(&order)
	return order, err.Error
}

func (o *orderRepository) GetOrder() ([]model.Order, error) {
	var orders []model.Order
	err := o.DB.Preload("Items").Find(&orders).Error
	return orders, err
}

func (o *orderRepository) UpdateOrder(ID int, order model.Order) error {
	return o.DB.Where("id=?", ID).Updates(&order).Error
}

func (o *orderRepository) DeleteORder(ID int) error {
	var order model.Order
	return o.DB.Where("id=?", ID).Delete(&order).Error
}
