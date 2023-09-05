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

func CreateInventory(in []*models.Inventory) (*mongo.InsertManyResult, error) {

	mcoll := config.GetCollection("inventory_SKU", "items")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.D{}
	result, err := mcoll.Find(ctx, filter, options.Find())
	var inventory []*models.Inventory_SKU
	// fmt.Println(result)
	if err != nil {
		fmt.Println("error1")

		fmt.Println(err.Error())
		return nil, err
	}
	for result.Next(ctx) {
		post := &models.Inventory_SKU{}
		err := result.Decode(post)
		if err != nil {
			fmt.Println("error2")

			return nil, err
		}
		inventory = append(inventory, post)
	}
	if err := result.Err(); err != nil {
		fmt.Println("error3")
		return nil, err
	}
	// fmt.Println(inventory)
	n := 0
	for j := 0; j < len(in); j++ {
		for i := n; i < n+10; i++ {
			in[j].Skus = append(in[j].Skus, inventory[i])
		}
		n = n + 10
	}
	fmt.Println("in", in)
	inv := []interface{}{}
	for v := 0; v < len(in); v++ {
		inv = append(inv, in[v])
	}
	// inv := []interface{}(in)
	res, err := InventoryContext().InsertMany(context.Background(), inv)
	if err != nil {
		fmt.Println("error4")
		return nil, err
	}
	// fmt.Println(res)
	return res, nil
}

func DeleteItems(item string, sku string, quantity float64) string {
	// filter := bson.D{{Key: "item", Value: item}, {Key: "skus.sku", Value: sku}}
	// update := bson.D{
	//     {Key: "$skus", Value: bson.D{
	//         {Key: "$elemMatch", Value: bson.D{
	//             {Key: "quantity", Value: bson.D{
	//                 {Key: "$inc", Value: -quantity},
	//             }},
	//         }},
	//     }},
	// }

	// // fv := bson.M{"$set": bson.M{"skus.quantity": quantity}}
	// _,err := InventoryContext().UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "failed"
	// }
	// return "success"
	filter := bson.D{
        {"item", item},
        {"skus.sku", sku}, // Match the specific SKU within the "skus" array by SKU name.
    }

    update := bson.D{
        {"$inc", bson.D{
            {"skus.$.quantity", -quantity}, // Decrement the "quantity" field by decrementAmount.
        }},
    }

    _, err := InventoryContext().UpdateOne(context.Background(), filter, update)
    if err != nil {
        return "failed"
    }

    return "success"

}
