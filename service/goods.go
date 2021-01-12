package service

import (
	"github.com/liuhongdi/digv16/dao"
	"github.com/liuhongdi/digv16/model"
)

//得到一件商品的详情
func GetOneGoods(goodsId int64) (*model.Goods, error) {
	return dao.SelectOneGoods(goodsId)
}

