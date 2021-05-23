package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sp/control"
)

func main() {

	// 禁用控制台颜色
	err:=initLogrus()
	if err != nil {
		panic(err)
	}
	router := gin.Default()

	obj := control.NewSp()
	v2 := router.Group("/v1")

	{
		v2.GET("/get_stock_prediction_data", obj.GetStockPredictionData)
	}

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

var log = logrus.New() // 创建一个log示例

func initLogrus() (err error) { // 初始化log的函数
	log.Formatter = &logrus.JSONFormatter{}                                       // 设置为json格式的日志
	f, err := os.OpenFile("./gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644) // 创建一个log日志文件
	if err != nil {
		return
	}
	log.Out = f                  // 设置log的默认文件输出
	gin.SetMode(gin.DebugMode) // 线上模式，控制台不会打印信息
	gin.DefaultWriter = io.MultiWriter(log.Out,os.Stdout)  // gin框架自己记录的日志也会输出

	log.Level = logrus.DebugLevel // 设置日志级别
	return
}
