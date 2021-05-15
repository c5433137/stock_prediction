package main

import (
	"github.com/gin-gonic/gin"
	"sp/control"
)



func main() {

	// Engin

	router := gin.Default()
	obj:= control.NewSp()
	v2 := router.Group("/v1")

	{
		v2.POST("/get_stock_prediction_data", obj.GetStockPredictionData)
	}

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}  
}
