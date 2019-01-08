package addressWatcher

import (
	"yunplanetBridge/models"
	"fmt"
	"github.com/astaxie/beego"
	"walletapi2/tool"
	ethClient2 "yunplanetBridge/ethClient"
	"yunplanetBridge/contracts/dexcontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"yunplanetBridge/models/jobModels"
)

type AddressWatcher struct {

}

func (this *AddressWatcher) Init() {
	this.querytrans()
}

func (this *AddressWatcher) querytrans() {

	var rpcClient ethClient2.EthRpcClient
	yuntclient,err:=rpcClient.GetyuntClient()
	if err!=nil{

	}

	job,err:=jobModels.GetJob("coinjob")
	if err!=nil&&err!=orm.ErrNoRows{

		logs.Info("",err.Error())
		return
	}

	tokenAddress :=common.HexToAddress("")
	fromAddress :=common.HexToAddress("")

	instance, err := dexcontracts.NewDexcontracts(tokenAddress, yuntclient)
	var results []models.Transact
	results=models.GetTokenTransList(int64(job.User_address_id),"yun",job.Limitcount)
	for _,temp := range results{
		fmt.Println("\r\n........>")
		address:=temp.Addressto
		toAddress :=common.HexToAddress(address)
		yuntaddress :=this.getyunt(address)
		fmt.Println(yuntaddress)
		value :=tool.FloatStringToBigInt(temp.Coinvalue2,8)
		instance.ExChangeYUN(&bind.TransactOpts{},fromAddress,toAddress,value)
	}
	endid :=job.User_address_id+job.Limitcount
	jobModels.UpdateJob("coinjob",endid)
}

func (this *AddressWatcher) getyunt(address string) (string) {
	var url string =beego.AppConfig.String("publicnode")+"/api/1.0/Account/Getaddress"
	b :="{\"address\":\""+address+"\"}"
	addressYunt:=tool.HttpPost(url,b)
	return  addressYunt
}
