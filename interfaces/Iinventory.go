package interfaces

import (
	"inventory/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Inventory interface {
	   GetInventoryByID(id int64) (*models.Inventory, error)
	   CreateInventory(i *models.Inventory) (*mongo.InsertOneResult, error)
}