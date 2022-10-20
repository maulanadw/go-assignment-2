package service

import (
	"go-assignment-2/model"
	"go-assignment-2/param"
	"go-assignment-2/repository"
)

type ItemService interface {
	CreateItem(orderID int, request []param.Item) error
	UpdateItemByOrderID(ID int, request []param.Item) error
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

func toItemModels(orderID int, itemParams []param.Item) []model.Item {
	itemModels := make([]model.Item, len(itemParams))
	for idx, itemParam := range itemParams {
		itemModels[idx] = toItemModel(orderID, itemParam)
	}

	return itemModels
}

func toItemModel(orderID int, itemParam param.Item) model.Item {
	return model.Item{
		ItemCode:    *itemParam.ItemCode,
		Description: *itemParam.Description,
		Quantitiy:   *itemParam.Quantity,
		OrderID:     orderID,
	}
}

func (i *itemService) CreateItem(orderID int, request []param.Item) error {
	itemModels := toItemModels(orderID, request)

	return i.itemRepository.CreateItem(itemModels)
}

func (i *itemService) UpdateItemByOrderID(ID int, request []param.Item) error {
	itemModels := toItemModels(ID, request)

	return i.itemRepository.UpdateItemByOrderID(ID, itemModels)
}

func (i *itemService) DeleteItem(ID int) error {
	panic("implement me")
}
