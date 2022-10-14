package service

import (
	"go-assignment-2/model"
	"go-assignment-2/param"
	"go-assignment-2/repository"
)

type ItemService interface {
	CreateItem(request param.Item, orderID int) error
	UpdateItem(ID int, request param.Item) error
	DeleteItem(ID int) error
}

type itemService struct {
	itemRepository repository.ItemRepository
}

func NewItemService(item repository.ItemRepository) ItemService {
	return &itemService{
		itemRepository: item,
	}
}

func toParamItems(itemModels []model.Item) []param.Item {
	itemParams := make([]param.Item, len(itemModels))
	for index, itemModel := range itemModels {
		itemParams[index] = toParamItem(itemModel)
	}

	return itemParams
}

func toParamItem(itemModel model.Item) param.Item {
	return param.Item{
		ItemCode:    &itemModel.ItemCode,
		Description: &itemModel.Description,
		Quantity:    &itemModel.Quantitiy,
	}
}

func (i *itemService) CreateItem(request param.Item, orderID int) error {
	itemModel := model.Item{
		ItemCode:    *request.ItemCode,
		Description: *request.Description,
		Quantitiy:   *request.Quantity,
		OrderID:     orderID,
	}

	return i.itemRepository.CreateItem(itemModel)
}

func (i *itemService) UpdateItem(ID int, request param.Item) error {
	var itemModel model.Item
	if request.ItemCode != nil {
		itemModel.ItemCode = *request.ItemCode
	}
	if request.Description != nil {
		itemModel.Description = *request.Description
	}
	if request.Quantity != nil {
		itemModel.Quantitiy = *request.Quantity
	}

	return i.itemRepository.UpdateItem(ID, itemModel)
}

func (i *itemService) DeleteItem(ID int) error {
	return i.itemRepository.DeleteItem(ID)
}
