package DTO

import "grocerystore/domain"

type ItemDTO struct {
	Id           int64
	Name         string
	PricePerUnit float64
	Quantity     int64
}

func NewItemDTO(i domain.ItemInterface) ItemDTO {
	return ItemDTO{
		Id:           i.GetId(),
		Name:         i.GetName(),
		PricePerUnit: i.GetPricePerUnit(),
		Quantity:     i.GetQuantity(),
	}
}
