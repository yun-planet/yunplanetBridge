package jobModels

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Address_checkjob struct {
	Id int32  `json:"id" orm:"column(id);pk;auto"`
	User_address_id int32  `orm:"column(user_address_id)" json:"user_address_id"`
	Limitcount int32  `orm:"column(limitcount)" json:"limitcount"`
	CreateTime time.Time  `orm:"column(create_time)" json:"create_time"`
	Jobtype string  `orm:"column(jobtype)" json:"jobtype"`
	Remark string  `orm:"column(remark)" json:"remark"`
}

func init() {
	orm.RegisterModel(new(Address_checkjob))
}


func GetJob(jobtype string)  (Address_checkjob,error){
	orm.Debug=true;
	o := orm.NewOrm()
	o.Using("jobdb")

	var job Address_checkjob

	trade := new(Address_checkjob)
	qs := o.QueryTable(trade).Filter("jobtype", jobtype)
	error:=qs.One(&job)
	return job,error
}

func UpdateJob(jobtype string,tableid int32)  (bool){

	orm.Debug=true
	o := orm.NewOrm()
	o.Using("jobdb")

	num, err := o.QueryTable("address_checkjob").Filter("jobtype", jobtype).Update(orm.Params{
		"user_address_id": tableid,
	})
	fmt.Println(err)
	return num>0

}