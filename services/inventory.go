package services

import (
	"context"
	"fmt"
	"inventory/config"
	"inventory/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type Invent struct {
// 	ctx             context.Context
// 	mongoCollection *mongo.Collection
// }

// func InitInventory(collection *mongo.Collection, ctx context.Context) interfaces.Inventory {
// 	return &Invent{ctx, collection}
// }

func InventoryContext() *mongo.Collection {
	return config.GetCollection("inventory_SKU", "inventory")
}

// func (i *Invent) GetInventoryByID(id int64) (*models.Inventory, error) {
// 	filter := bson.D{{Key: "_id", Value: id}}
// 	var inventory *models.Inventory
// 	res := i.mongoCollection.FindOne(i.ctx, filter)
// 	err := res.Decode(&inventory)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return inventory, nil
// }

func CreateInventory(in *models.Inventory) (*mongo.InsertOneResult, error) {

	mcoll := config.GetCollection("inventory_SKU", "items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{}
	result, err := mcoll.Find(ctx, filter, options.Find().SetLimit(10))
	var inventory []*models.Sku
	// fmt.Println(result)
	if err != nil {
		fmt.Println("error")

		fmt.Println(err.Error())
		return nil, err
	}
	for result.Next(ctx) {
		post := &models.Sku{}
		err := result.Decode(post)
		if err != nil {
			fmt.Println("error")

			return nil, err
		}
		inventory = append(inventory, post)
	}
	if err := result.Err(); err != nil {
		fmt.Println("error")
		return nil, err
	}
	// fmt.Println(inventory)
	in.Skus = inventory
	fmt.Println("in", in)
	res, err := InventoryContext().InsertOne(context.Background(), in)
	if err != nil {
		return nil, err
	}
	// fmt.Println(res)
	return res, nil
}
