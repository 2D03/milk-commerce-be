package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ProductModel = DBModel{
	ColName:  "products",
	Template: Product{},
}

type Product struct {
	CreatedTime     *time.Time `json:"createdTime" bson:"created_time,omitempty"`
	LastUpdatedTime *time.Time `json:"lastUpdatedTime" bson:"last_updated_time,omitempty"`
	ProductId       string     `json:"productId,omitempty" bson:"product_id,omitempty"`
	Name            string     `json:"name,omitempty" bson:"name,omitempty"`
	Category        string     `json:"category,omitempty" bson:"category,omitempty"`
	Price           int64      `json:"price,omitempty" bson:"price,omitempty"`
	Amount          int        `json:"amount,omitempty" bson:"amount,omitempty"`
	Description     string     `json:"description,omitempty" bson:"description,omitempty"`
}

func InitProductModel(client *mongo.Client, ctx context.Context, dbName string) error {
	ProductModel.DBName = dbName
	err := ProductModel.Init(client, ctx)
	if err != nil {
		panic(err)
	}

	return nil
}
