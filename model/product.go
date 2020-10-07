package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

var ProductModel = DBModel{
	ColName:  "products",
	Template: Product{},
}

type Product struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
}

func InitProductModel(client *mongo.Client, ctx context.Context, dbName string) error {
	ProductModel.DBName = dbName
	err := ProductModel.Init(client, ctx)
	if err != nil {
		panic(err)
	}

	return nil
}
