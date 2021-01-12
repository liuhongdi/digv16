package service

import (
	"github.com/liuhongdi/digv16/dao"
	"github.com/liuhongdi/digv16/model"
)

//新加一个订单
func AddOneOrder(goodsId int64,buyNum int) (*model.Order, error) {
	return dao.AddOneOrder(goodsId,buyNum)
}

//新加一个订单,使用tx
func AddOneOrderTx(goodsId int64,buyNum int) (*model.Order, error) {
	return dao.AddOneOrderTx(goodsId,buyNum)
}