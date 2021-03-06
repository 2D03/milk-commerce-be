package api

import (
	"context"
	"fmt"
	"github.com/Kajekk/milk-commerce-be/model"
	"github.com/gin-gonic/gin"
	"time"
)

func CreateProducts(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	//item := model.Product{
	//	Name: "qwe",
	//}

	myItem, err := model.ConvertToBson(product)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	myItem["created_time"] = time.Now()
	myItem["last_updated_time"] = time.Now()
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
