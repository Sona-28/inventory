package main

import (
	"context"
	"fmt"
	h "inventory/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:2023", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect: ", err)
	}
	defer conn.Close()
	client := h.NewInventoryServiceClient(conn)
	// req := &h.AllInventoryItems{
	// 	Items: []*h.InventoryItem{
	// 		{
	// 			Id:   1,
	// 			Item: "shirt",
	// 			Features: []string{
	// 				"slim fit",
	// 				"full sleeves",
	// 			},
	// 			Categories: []string{
	// 				"tshirt",
	// 				"formal",
	// 			},
	// 		},
	// 		{Id: 2,
	// 			Item: "pants",
	// 			Features: []string{
	// 				"full length",
	// 				"slim fit",
	// 			},
	// 			Categories: []string{
	// 				"jeans",
	// 				"formal",
	// 			},
	// 		},
	// 		{Id: 3,
	// 			Item: "headphones",
	// 			Features: []string{
	// 				"wireless",
	// 				"noise cancellation",
	// 			},
	// 			Categories: []string{
	// 				"headphones",
	// 				"bluetooth",
	// 			},
	// 		},
	// 	},
	// }
	// response, err := client.CreateInventory(context.Background(), req)
	response, err := client.UpdateInventory(context.Background(), &h.ItemToDelete{
		Item:     "shirt",
		Sku:      "SKU002",
		Quantity: 10,
	})
	if err != nil {
		log.Fatal("Failed to call: ", err)
	}
	fmt.Println("Response: ", response)
}
