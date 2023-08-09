package main

import (
	"fmt"
	"ms/controller"
	"ms/model"
	"net/http"
)
var successNum int32
var failNum int32
func init(){
	fmt.Println("开始初始化")
	fmt.Println("开始初始化")
	fmt.Println("开始初始化")
	fmt.Println("开始初始化")
	err := model.NewClient()
	if err != nil {
		fmt.Print("数据库连接失败"+err.Error())
		return
	}
	fmt.Println("数据库连接成功")
	_ ,err = controller.InitGoods()
	if err != nil {
		fmt.Print("初始化商品失败"+err.Error())
		return
	}
	fmt.Println("初始化商品成功")
}
func main(){
	http.HandleFunc("/buy/ticket", handleReq)
	http.ListenAndServe(":3005", nil)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	goodsName := "苹果"
	err := controller.SoldGoods(goodsName)
	if err != nil {
		failNum += 1
		str := fmt.Sprintf("%d购买%v失败%v",failNum,goodsName,err.Error())
		fmt.Println(str)

		return
	}
	successNum += 1
	str := fmt.Sprintf("%d购买%v成功",successNum,goodsName)
	fmt.Println(str)
}

