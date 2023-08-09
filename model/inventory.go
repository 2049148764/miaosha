package model

type Inventory struct {
	GoodsName   string
	StockNum  int
	SoldNum   int
	GoodsNum   int
}
func GetStockNum(i *Inventory)(stockNum int,soldNum int , err error){
	//defer DB.Close()

	res := DB.QueryRow("select stockNum , soldNum from goods where goodsName = ?",i.GoodsName)
	if err != nil {
		return
	}
	err = res.Scan(&i.StockNum,&i.SoldNum)
	//fmt.Println(err.Error())
	if err != nil {
		return
	}
	stockNum = i.StockNum
	soldNum = i.SoldNum
	return
}

func SoldGoods(i *Inventory)(rows int64 ,err error){
	//defer DB.Close()
	res,err := DB.Exec("update goods set stockNum= ? , soldNum =  ? where goodsName = ?",i.StockNum,i.SoldNum,i.GoodsName)
	if err != nil {
		return
	}
	rows , err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}

func InitGoods(i *Inventory)(rows int64,err error){
	//defer DB.Close()
	res := DB.QueryRow("select count(id) from goods where goodsName = ?",i.GoodsName)
	if err != nil {
		return
	}
	var count int
	err = res.Scan(&count)
	if err != nil || count > 0 {
		return
	}
	res1,err := DB.Exec("insert INTO goods(goodsName,stockNum,soldNum,goodsNum) values(?,?,?,?)",i.GoodsName,i.StockNum,i.SoldNum,i.GoodsNum)
	if err != nil {
		return
	}
	rows,err = res1.RowsAffected()
	if err != nil {
		return
	}
	return
}
