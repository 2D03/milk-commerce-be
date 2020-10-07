package api

import (
	"context"
	"fmt"
	"github.com/Kajekk/milk-commerce-be/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func CreateProducts(c *gin.Context) {
	item := model.Product{
		Name: "asd",
	}
	obj, err := bson.Marshal(item)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	myBson := bson.M{}
	_ = bson.Unmarshal(obj, &myBson)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	time.Sleep(time.Duration(5) * time.Second)
	rs, err := model.ProductModel.Collection.InsertOne(ctx, myBson)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if rs != nil {
		c.JSON(200, gin.H{
			"message": "created",
		})
		fmt.Println(rs)
		return
	}
}
