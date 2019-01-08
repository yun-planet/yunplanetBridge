package ethClient

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthRpcClient struct {

}

func (this *EthRpcClient) GetEtherClient() (*ethclient.Client,error){

	web3Url:=beego.AppConfig.String("yunnode")
	fmt.Println(web3Url)

	client, err := ethclient.Dial(web3Url)
	fmt.Println("\r\n client:")
	fmt.Println(err)
	return client,err
}

func (this *EthRpcClient) GetyuntClient() (*ethclient.Client,error){

	web3Url:=beego.AppConfig.String("yuntnode")
	fmt.Println(web3Url)

	client, err := ethclient.Dial(web3Url)
	fmt.Println("\r\n client:")
	fmt.Println(err)
	return client,err
}