package control

import (
	"github.com/gin-gonic/gin"
	"sp/logic"
	"strconv"
	"strings"
)

type Sp struct {

}

func NewSp() *Sp {
	return &Sp{}
}

//GetStockPredictionData 获取股票预测数据
func (s *Sp) GetStockPredictionData(g *gin.Context)  {

	stockId:=g.Query("stockId")
	stockId = strings.ToLower(stockId)
	stockId = strings.Replace(stockId," ","",-1)
	if stockId == ""{
		stockId="002156"
	}
	market:=g.Query("market")
	market = strings.ToLower(market)
	if market == ""{
		market=strings.ToLower("sz")
	}
	iv:=g.Query("iv")
	_iv,_:=strconv.ParseFloat(iv,64)
	pv:=g.Query("pv")
	_pv,_:=strconv.ParseFloat(pv,64)
	ov:=g.Query("ov")
	_ov,_:=strconv.ParseFloat(ov,64)
	var data []logic.CodeListNewType
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
	res.Stock = market+stockId
	res = logic.Prediction(data,_iv,_pv,_ov)
	g.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"data":res,
	})

}
