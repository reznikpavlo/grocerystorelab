package domain

import (
	"strconv"
	"strings"
)

type Item struct {
	id           int64
	name         string
	pricePerUnit float64
	quantity     int64
}

func NewItem() *Item {
	return &Item{}

}
func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetPricePerUnit() float64 {
	return i.pricePerUnit
}

func (i *Item) GetQuantity() int64 {
	return i.quantity
}

func (i *Item) SetQuantity(q int64) *Item {
	i.quantity = q
	return i
}

func (i *Item) SetName(n string) *Item {
	i.name = n
	return i
}

func (i *Item) SetPricePerUnit(p float64) *Item {
	i.pricePerUnit = p
	return i
}

func (i *Item) GetId() int64 {
	return i.id
}

func (i *Item) SetId(idd int64) {
	i.id = idd
}

func (i *Item) String() string {
	builder := strings.Builder{}
	builder.WriteString("{name: ")
	builder.WriteString(i.name)
	builder.WriteString(" quantity: ")
	builder.WriteString(strconv.FormatInt(i.quantity, 10))
	builder.WriteString(" price: ")
	builder.WriteString(strconv.FormatFloat(i.pricePerUnit, 'f', 2, 64))
	builder.WriteString("}")
	return builder.String()
}
