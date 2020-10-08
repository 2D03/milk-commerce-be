package api

import (
	"context"
	"fmt"
	"github.com/Kajekk/milk-commerce-be/model"
	"github.com/gin-gonic/gin"
	"time"
)

func UpdateProduct(c *gin.Context) {
	//TODO Update Products
	item := model.Product{
		Name: "asd",
	}
	myItem, err := model.ConvertToBson(item)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//time.Sleep(time.Duration(5) * time.Second)
	rs, err := model.ProductModel.Collection.InsertOne(ctx, myItem)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if rs != nil {
		c.JSON(200, model.APIResponse{
			Status:  model.APIStatus.Ok,
			Message: "Successfully",
		})
		fmt.Println(rs)
		return
	}
}
