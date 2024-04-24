package service

import (
	"errors"
	"grocerystore/domain"
)

type ItemService interface {
	GetItems() map[int64]domain.ItemInterface
	SellItem(id int64) (bool, error)
}

type itemServiceImplementation struct {
	store domain.Store
}

func NewItemService() *itemServiceImplementation {
	item := domain.NewItem().
		SetName("cherry").
		SetQuantity(10).
		SetPricePerUnit(5)

	item2 := domain.NewItem().
		SetName("tomato").
		SetQuantity(10).
		SetPricePerUnit(3)
	item3 := domain.NewItem().
		SetName("potato").
		SetQuantity(10).
		SetPricePerUnit(2)
	itemsMap := map[int64]domain.ItemInterface{
		1: item,
		2: item2,
		3: item3,
	}
	store := domain.Store{ItemsMap: itemsMap}
	return &itemServiceImplementation{store: store}
}

func (i *itemServiceImplementation) GetItems() map[int64]domain.ItemInterface {
	return i.store.ItemsMap
}

func (i itemServiceImplementation) SellItem(id int64) (bool, error) {
	return false, errors.New("sellItem in ItemService is not implemented")
}
