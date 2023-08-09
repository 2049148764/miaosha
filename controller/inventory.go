package controller

import (
	"errors"
	"fmt"
	model "ms/model"
)

type A struct {

}
var a chan A
//InitGoods
func InitGoods()(rows int64,err error){
	i := model.Inventory{}
	i.GoodsName = "苹果"
	i.GoodsNum = 100000
	i.StockNum = 100000
	i.SoldNum = 0
	rows ,err = model.InitGoods(&i)
	if err != nil {
		return
	}
	a = make(chan A, 1)
	a<- A{}
	return
}

func SoldGoods(goodsName string)(err error){
    <-a
	i := model.Inventory{}
	i.GoodsName = goodsName
	stockNum,soldNum,err := model.GetStockNum(&i)
	if err != nil {
		err = errors.New("获取库存失败,原因是----"+err.Error())
		a<- A{}
		return
	}
	fmt.Println(stockNum)
	if stockNum <= 0 {
		err =  errors.New("库存不够"+ string(stockNum))
		a<- A{}
		return
	}
	i.StockNum = stockNum - 1
    i.SoldNum = soldNum + 1
    _ ,err = model.SoldGoods(&i)
    if err != nil {
		a<- A{}
    	return
	}
	a<- A{}
	return
}