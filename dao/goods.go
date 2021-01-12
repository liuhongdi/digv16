package dao

import (
	"fmt"
	"github.com/liuhongdi/digv16/global"
	"github.com/liuhongdi/digv16/model"
)

//得到一件商品的详情
func SelectOneGoods(goodsId int64) (*model.Goods, error) {
	goodsOne:=&model.Goods{}
	err := global.DBLink.Where("goodsId=?",goodsId).First(&goodsOne).Error
	fmt.Println(goodsOne)
	if (err != nil) {
		return nil,err
	} else {
		return goodsOne,nil
	}
}
