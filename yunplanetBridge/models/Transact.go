package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Transact struct {
	Id int64 `json:"id"`
	Addressto string `orm:"column(addressTo)";json:"addressTo"`
	Blocknumber string `orm:"column(blockNumber)";json:"blockNumber"`
	Txthash string `orm:"column(txtHash)";json:"txtHash"`
	Cointype string `orm:"column(coinType)";json:"coinType"`
	Coinvalue string `orm:"column(coinValue)";json:"coinValue"`
	Create_time string `orm:"column(create_time)";json:"create_time"`
	Transtype string `orm:"column(transType)";json:"transType"`
	Coinnumber string `orm:"column(coinNumber)";json:"coinNumber"`
	Addressfrom string `orm:"column(addressFrom)";json:"addressFrom"`
	Handerstatus int `orm:"column(handerStatus)";json:"handerStatus"`
	Transtatus int `orm:"column(transtatus)";json:"tranStatus"`
	Coinvalue2 string `orm:"column(coinValue2)";json:"coinValue2"`
	Nonce int64 `orm:"column(nonce)";json:"nonce"`
}


func init(){
	orm.RegisterModel(new(Transact))
}

func GetTokenTransList(id int64,coinType string, limit int32)  []Transact{
	orm.Debug=true
	o := orm.NewOrm()
	o.Using("default")
	var transactList []Transact
	sql:=fmt.Sprintf("SELECT * FROM transact  WHERE id> %d  and coinType ='%s' and transtatus=1 and handerStatus=0 order by id  limit %d ",id,coinType, limit)

	_,err := o.Raw(sql).QueryRows(&transactList)
	if err !=nil{
		fmt.Println(err)
	}

	return transactList
}
