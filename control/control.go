package control

import (
	"github.com/gin-gonic/gin"
	"sp/logic"
	"strings"
)

type Sp struct {

}

func NewSp() *Sp {
	return &Sp{}
}

//GetStockPredictionData 获取股票预测数据
func (s *Sp) GetStockPredictionData(g *gin.Context)  {

	market:=strings.ToLower("sz")
	stockId:="002156"
	var data []logic.CodeList
	switch market {
	case "sz","sh":
		data=logic.GetStockData(stockId)
	default:
		g.JSON(400,gin.H{
			"code":400,
			"msg":"参数错误",
		})
		return
	}
	var res logic.ResData
	res = logic.Prediction(data)
	g.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"data":res,
	})

}
