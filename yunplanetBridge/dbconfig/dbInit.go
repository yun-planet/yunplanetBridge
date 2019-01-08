package dbconfig

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type DBInit struct {

}

func (this *DBInit) Init() {
	yunpalnetdbhost := "127.0.0.1"
	yunpalnetdbport := "3606"
	yunpalnetdbuser := "user"
	yunpalnetdbpassword := "yourpass"
	yunpalnetdb := "yourdb"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	conn := yunpalnetdbuser + ":" + yunpalnetdbpassword + "@tcp(" + yunpalnetdbhost + ":" + yunpalnetdbport + ")/" + yunpalnetdb + "?charset=utf8"
	fmt.Println(conn)
	orm.RegisterDataBase("default", "mysql", conn,30,300)
}