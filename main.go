package main

import (
	"context"
	"fmt"
	"github.com/Kajekk/milk-commerce-be/api"
	"github.com/Kajekk/milk-commerce-be/conf"
	"github.com/Kajekk/milk-commerce-be/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

var client *mongo.Client
var ctx context.Context

func main() {
	env := os.Getenv("env")
	version := os.Getenv("version")
	config := os.Getenv("config")

	setupDB()
	defer client.Disconnect(ctx)
	//collection := client.Database("testing").Collection("numbers")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"env":     env,
			"config":  config,
			"version": version,
		})
	})
	router.GET("/product", api.GetProducts)
	router.POST("/product", api.CreateProducts)
	router.PUT("/product", api.UpdateProduct)
	router.DELETE("/product", api.RemoveProduct)

	_ = router.Run(":8000")

}

func setupDB() {
	println("Connecting db")
	//configMap, err := conf.GetConfigDB()
	//if err != nil {
	//	panic(err)
	//}

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	//databases, err := client.ListDatabaseNames(ctx, bson.M{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(databases)
	onDBConnected(client, ctx)

	fmt.Println("Successfully connected and pinged.")
}

func onDBConnected(client *mongo.Client, ctx context.Context) {
	_ = model.InitProductModel(client, ctx, conf.Config.MainDBName)
}
