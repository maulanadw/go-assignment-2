package service

import (
	"go-assignment-2/model"
	"go-assignment-2/param"
	"go-assignment-2/repository"
	"time"
)

type OrderService interface {
	CreateOrder(customerName string) error
	GetOrder() ([]param.Order, error)
	UpdateOrder(ID int, request param.Order) error
	DeleteOrder(ID int) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(order repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: order,
	}
}

func toOrderParams(orderModels []model.Order) []param.Order {
	orderParams := make([]param.Order, len(orderModels))
	for index, orderModel := range orderModels {
		orderParams[index] = toOrderParam(orderModel)
	}

	return orderParams
}

func toOrderParam(orderModel model.Order) param.Order {
	return param.Order{
		CustomerName: &orderModel.CustomerName,
		Items:        toParamItems(orderModel.Items),
		OrderedAt:    &orderModel.OrderedAt,
	}
}

func (o *orderService) CreateOrder(customerName string) error {
	order := model.Order{
		OrderedAt:    time.Now(),
		CustomerName: customerName,
	}

	return o.orderRepository.CreateOrder(order)
}

func (o *orderService) GetOrder() ([]param.Order, error) {
	orders, err := o.orderRepository.GetOrder()
	if err != nil {
		return nil, err
	}

	orderParams := toOrderParams(orders)
	return orderParams, nil
}

func (o *orderService) UpdateOrder(ID int, request param.Order) error {
	var order model.Order
	if request.CustomerName != nil {
		order.CustomerName = *request.CustomerName
	}
	if request.OrderedAt != nil {
		order.OrderedAt = *request.OrderedAt
	}

	return o.orderRepository.UpdateOrder(ID, order)
}

func (o *orderService) DeleteOrder(ID int) error {
	return o.orderRepository.DeleteORder(ID)
}
