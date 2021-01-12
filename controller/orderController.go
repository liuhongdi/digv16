package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv16/global"
	"github.com/liuhongdi/digv16/service"
)

type OrderController struct{}

//result:= global.NewResult(c)

func NewOrderController() OrderController {
	return OrderController{}
}

//新加一个订单
func (o *OrderController) AddOne(c *gin.Context) {
	result := global.NewResult(c)

	goodsId:=1;
	buyNum:=11;

	orderOne,err := service.AddOneOrder(int64(goodsId),buyNum);
	if err != nil {
		result.Error(404,"数据处理错误")
	} else {
		result.Success(&orderOne);
	}
	return
}


//新加一个订单,使用tx
func (o *OrderController) AddOneTx(c *gin.Context) {
	result := global.NewResult(c)
	goodsId:=1;
	buyNum:=11;

	orderOne,err := service.AddOneOrderTx(int64(goodsId),buyNum);
	if err != nil {
		result.Error(404,"数据处理错误")
	} else {
		result.Success(&orderOne);
	}
	return
}