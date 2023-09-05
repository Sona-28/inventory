package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
	ID         int64    `json:"id" bson:"_id"`
	Item       string   `json:"item" bson:"item"`
	Features   []string `json:"features" bson:"features"`
	Categories []string `json:"categories" bson:"categories"`
	Skus       []Inventory_SKU   `json:"skus" bson:"skus"`
}

type Inventory_SKU struct {
	Id 	 primitive.ObjectID `json:"_id" bson:"_id"`
	Sku      string       `json:"sku" bson:"sku"`
	Price    Price_type   `json:"price" bson:"price"`
	Quantity float64      `json:"quantity" bson:"quantity"`
	Options  Options_type `json:"options" bson:"options"`
}

type Price_type struct {
	Base     float64 `json:"base" bson:"base"`
	Currency string  `json:"currency" bson:"currency"`
	Discount float64 `json:"discount" bson:"discount"`
}
type Options_type struct {
	Size     Size_type `json:"size" bson:"size"`
	Features []string  `json:"features" bson:"features"`
	Colors   []string  `json:"colors" bson:"colors"`
	Ruling   string    `json:"ruling" bson:"ruling"`
	Image    string    `json:"image" bson:"image"`
}

type Size_type struct {
	H float64 `json:"h" bson:"h"`
	L float64 `json:"l" bson:"l"`
	W float64 `json:"w" bson:"w"`
}
