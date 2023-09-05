package main

import (
	"fmt"
	"inventory/interfaces"
	// "inventory/models"
	"inventory/services"
)

var IService interfaces.Inventory

func main(){
	// client, _ := config.ConnectDataBase()
	// mcoll := client.Database("inventory_SKU").Collection("inventory")
	// services.InitInventory(mcoll, context.Background())
	// inventory := []*models.Inventory{
	// 	{ID:         1,
	// 		Item:       "shirt",
	// 		Features:   []string{
	// 			"slim fit",
	// 			"full sleeves",
	// 		},
	// 		Categories: []string{
	// 			"tshirt",
	// 			"formal",
	// 		},
	// 	},
	// 	{ID:         2,
	// 		Item:       "pants",
	// 		Features:   []string{
	// 			"full length",
	// 			"slim fit",
	// 		},
	// 		Categories: []string{
	// 			"jeans",
	// 			"formal",
	// 		},
	// 	},
	// 	{ID:         3,
	// 		Item:       "headphones",
	// 		Features:   []string{
	// 			"wireless",
	// 			"noise cancellation",
	// 		},
	// 		Categories: []string{
	// 			"headphones",
	// 			"bluetooth",
	// 		},
	// 	},
	// }
	// res, err:= services.CreateInventory(inventory)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(res)
	res := services.DeleteItems("shirts","SKU002", 2.0)
	fmt.Println(res)
}