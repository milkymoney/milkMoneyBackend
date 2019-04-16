package models

import(

)

func init(){

}


/*接受关系为用户到task的多对多关系，一个用户可以接受多个任务，一个任务也可以被多个用户所接受
但是实现上为一对一关系，即用户id与任务id绑定，一起组成acId
原则上一个用户不能够接受*/
type AcceptRelation struct {
	Id			int 		`orm:"pk"`//自增数字序列作为主键
	AcceptDate	string
	User 		*User 		`orm:"reverse(many)"`
	Task		*Task 		`orm:"reverse(many)"`
}

type ReleaseRelation struct{
	Id			int 		`orm:"pk"`//自增数字序列作为主键
	ReleaseDate	string
	User 		*User 		`orm:"reverse(many)"`
	Task		*Task 		`orm:"reverse(many)"`
}
