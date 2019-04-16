package models

import (
	"github.com/astaxie/beego/orm"

)

func init(){

}
//注：原则上，所有的输入数据合法性检查由controller处进行。在models进行的是与模型之间数据关系有关的检查，比如任务的最多接纳人数是否达到上界，等。
type Task struct{
	Id				string				`orm:"pk"`
	Type			string
	Description		string
	Reward			float32
	Deadline 		string
	Label			[]string
	State			string
	Priority		int
	MaxAccept		int //任务同时允许的最大接受人数
	AcceptRelation 	[]*AcceptRelation	`orm:"rel(fk)"`
	ReleaseRelation []*ReleaseRelation	`orm:"rel(fk)"`}

/*
函数目的：创建任务
调用时机：主要由controllers/task.go调用，实际上可用于任何时刻的添加
需要执行的任务：
	1.向数据库中写入数据

调用成功：直接返回taskId与nil
调用失败：返回""与err
*/
func CreateTask(task Task) (taskId string,err error){

}

/*
函数目的：获取任务
调用时机：任何需要通过taskId获取任务的场景
需要执行的任务：
	1.根据给定的taskId获取task对象

调用成功：直接返回task指针与nil
调用失败：返回nil与错误对象
*/

func GetTask(taskId string) (task *Task,err error){

}

/*
函数目的：更新任务信息
*/

func UpdateUser(taskId string,tt *Task) (task *Task,err error){

}