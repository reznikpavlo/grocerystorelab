package main

import (
	"encoding/json"
	"fmt"
	"grocerystore/domain/DTO"
	"grocerystore/service"
)

func main() {
	defer finish()

	service := service.NewItemService()
	items := service.GetItems()

	var itemsDTO map[int64]DTO.ItemDTO
	itemsDTO = make(map[int64]DTO.ItemDTO)
	for k, v := range items {
		itemsDTO[k] = DTO.NewItemDTO(v)
	}
	fmt.Println(items)
	itemsJSON, _ := json.Marshal(itemsDTO)
	fmt.Printf("%s\n", itemsJSON)

}

func finish() {
	if r := recover(); r != nil {
		fmt.Println("ERROR IN ", r)
	}
}
