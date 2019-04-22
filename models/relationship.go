package models

import(
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init(){

}


/*接受关系为用户到task的多对多关系，一个用户可以接受多个任务，一个任务也可以被多个用户所接受
但是实现上为一对一关系，即用户id与任务id绑定，一起组成acId
原则上一个用户不能够接受*/
type AcceptRelation struct {
	Id			int
	AcceptDate	string
	User		*User	`orm:"rel(fk)"`
	Task		*Task	`orm:"rel(fk)"`
}

type ReleaseRelation struct{
	Id			int
	ReleaseDate	string
	User		*User	`orm:"rel(fk)"`
	Task		*Task	`orm:"rel(fk)"`
}

/*业务函数群*/

//根据用户id和任务id创建新的AcceptRelation，并加入到数据库之中
func CreateNewAcRelById(userId,taskId int,acceptDate string) (*AcceptRelation,error){
	user,err := GetUserById(userId)
	if err != nil{
		return nil,err
	}
	task,err := GetTaskById(taskId)
	if err != nil{
		return nil,err
	}
	newRelation := &AcceptRelation{AcceptDate:acceptDate,User:user,Task:task}
	acId,err := CreateAcceptRelation(newRelation)
	if err != nil{
		return nil,err
	}
	newRelation.Id = acId
	return newRelation,nil
}

/*
函数目的：创建AcceptRelation
调用时机：需要将relation加入到数据库中
需要执行的任务：
	1.创建对象，并加入到数据库中

调用成功：返回这个对象的id,nil
调用失败："",err对象
	调用失败场景：暂时没有想到
*/

func CreateAcceptRelation(relation *AcceptRelation) (acId int,err error){
	o := orm.NewOrm()
	id64,err := o.Insert(relation)
	id := int(id64)
	if err != nil{
		return 0,err
	}
	return id,nil
}

/*
函数目的：拿到AcceptRelation
调用时机：需要使用userId和taskId拿到relation
需要执行的任务：
	1.从数据库中查询并返回对象

调用成功：返回这个对象,nil
调用失败：nil,err对象
	调用失败场景：查询不到对应的对象
*/
func GetAcceptRelation(userId,taskId int) (relation []*AcceptRelation,err error){
	var relations []*AcceptRelation
	o := orm.NewOrm()

	if _,err := o.QueryTable("accept_relation").Filter("user_id",userId).Filter("task_id",taskId).All(&relations); err != nil{
		return nil,fmt.Errorf("User id or task id not correct.")
	} else{
		return relations,nil
	}

}



/*
函数目的：删除AcceptRelation
调用时机：需要将relation从数据库中删除
需要执行的任务：
	1.删除对象，并加入到数据库中
	2.是否要拿到对应的user和task并从他们那里删除（我感觉不用，等待外键实验结果）

调用成功：返回这个对象,nil
调用失败：nil,err对象
	调用失败场景：查询不到对象

func DeleteAcceptRelation(userId,taskId int) error{

}

/*
函数目的：拿到ReleaseRelation
调用时机：需要使用userId和taskId拿到relation
需要执行的任务：
	1.从数据库中查询并返回对象

调用成功：返回这个对象,nil
调用失败：nil,err对象
	调用失败场景：查询不到对应的对象

func GetReleaseRelation(userId,taskId int) (relation *ReleaseRelation,err error){

}

/*
函数目的：创建ReleaseRelation
调用时机：需要将relation加入到数据库中
需要执行的任务：
	1.创建对象，并加入到数据库中

调用成功：返回这个对象的id,nil
调用失败："",err对象
	调用失败场景：暂时没有想到

func CreateReleaseRelation(relation *ReleaseRelation) (reId int,err error){

}

/*
函数目的：删除ReleaseRelation
调用时机：需要将relation从数据库中删除
需要执行的任务：
	1.删除对象，并加入到数据库中
	2.是否要拿到对应的user和task并从他们那里删除（我感觉不用，等待外键实验结果）

调用成功：返回这个对象,nil
调用失败：nil,err对象
	调用失败场景：查询不到对象

func DeleteReleaseRelation(userId,taskId int) error{

}
*/