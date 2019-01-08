package main

import (
	"fmt"
	_ "yunplanetBridge/routers"

	"github.com/astaxie/beego"
	"github.com/jasonlvhit/gocron"
	addressWatcher2 "yunplanetBridge/addressWatcher"
	"yunplanetBridge/dbconfig"
)

func init()  {
	dbInit:=dbconfig.DBInit{}
	dbInit.Init()
}

func main() {

	watcherInit()
	beego.Run()
}

func watcherInit() {
	s := gocron.NewScheduler()
	addressWatcher :=addressWatcher2.AddressWatcher{}
	s.Every(5).Seconds().Do(task,addressWatcher)
	sc := s.Start() // keep the channel
	go cotrollers(s, sc)  // wait
	<-sc
}

func task(jobObj addressWatcher2.AddressWatcher)  {
	fmt.Println("watcher trans")
	jobObj.Init()
}

func cotrollers(s *gocron.Scheduler, sc chan bool) {

}
