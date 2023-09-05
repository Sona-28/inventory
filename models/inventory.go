package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Inventory struct {
   ID       int64   `json:"id" bson:"_id"`
   Item       string    `json:"item" bson:"item"`
   Features    []string `json:"features" bson:"features"`
   Categories []string `json:"categories" bson:"categories"`
   Skus       []*Sku   `json:"skus" bson:"skus"`
}

type Sku struct {
	Id primitive.ObjectID `json:"id" bson:"_id"`
   Sku string `json:"sku" bson:"sku"`
   Price Price `json:"price" bson:"price"`
   Quantity int64 `json:"quantity" bson:"quantity"`
   Options Options `json:"options" bson:"options"`
}

type Price struct {
   Base float64 `json:"base" bson:"base"`
   Currency string `json:"currency" bson:"currency"`
   Discount float64 `json:"discount" bson:"discount"`
}

type Options struct {
   Size Size `json:"size" bson:"size"`
   Features []string `json:"features" bson:"features"`
   Colors []string `json:"colors" bson:"colors"`
   Ruling string `json:"ruling" bson:"ruling"`
   Image string `json:"image" bson:"image"`
}

type Size struct {
   Height float64 `json:"height" bson:"height"`
   Length float64 `json:"length" bson:"length"`
   Breadth float64 `json:"breadth" bson:"breadth"`
}
