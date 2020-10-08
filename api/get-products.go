package api

import (
	"context"
	"fmt"
	"github.com/Kajekk/milk-commerce-be/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func GetProducts(c *gin.Context) {
	productId := c.Query("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, &model.APIResponse{
			Status:    model.APIStatus.Invalid,
			Message:   "missing product",
			ErrorCode: "MISSING_PRODUCT_ID",
		})
		return
	}

	//item := model.Product{
	//	ProductId: productId,
	//}
	//myItem, err := model.ConvertToBson(item)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": err.Error(),
	//	})
	//	return
	//}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//time.Sleep(time.Duration(5) * time.Second)
	rs, err := model.ProductModel.Collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, &model.APIResponse{
			Status:  model.APIStatus.Error,
			Message: err.Error(),
		})
		return
	}
	var products []bson.M
	if err = rs.All(ctx, &products); err != nil {

	}
	fmt.Println(products)

	if rs != nil {
		c.JSON(http.StatusOK, &model.APIResponse{
			Status:  model.APIStatus.Ok,
			Message: "Successfully",
			Data:    products,
			Total:   int64(len(products)),
		})
		return
	}
}
