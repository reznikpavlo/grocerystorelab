package domain

type ItemInterface interface {
	GetId() int64
	SetId(idd int64)
	GetName() string
	GetPricePerUnit() float64
	GetQuantity() int64
	SetQuantity(q int64) *Item
	SetName(n string) *Item
	SetPricePerUnit(p float64) *Item
}
