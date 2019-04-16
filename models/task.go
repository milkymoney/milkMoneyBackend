package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init(){

}

type Task struct{
	Id				string			`orm:"pk"`
	Type			string
	Description		string
	Reward			float32
	Deadline 		time.Time		`orm:"type(datetime)"`
	Label			[]string
	State			string
	Priority		int
	MaxAccept		int //任务同时允许的最大接受人数
	AcceptRelation 	*AcceptRelation `orm:"reverse(one)"`
	ReleaseRelation *ReleaseRelation `orm:"reverse(one)"`
}