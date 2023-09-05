package main

import (
	"fmt"
	"inventory/interfaces"
	"inventory/models"
	"inventory/services"
)

var IService interfaces.Inventory

func main(){
	// client, _ := config.ConnectDataBase()
	// mcoll := client.Database("inventory_SKU").Collection("inventory")
	// services.InitInventory(mcoll, context.Background())
	inventory := &models.Inventory{
		ID:         1,
		Item:       "shirt",
		Features:   []string{
			"slim fit",
			"full sleeves",
		},
		Categories: []string{
			"tshirt",
			"formal",
		},
	}
	res, err:= services.CreateInventory(inventory)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}