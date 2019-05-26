package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init(){

}

type TaskState string

const(
	Task_pend 	TaskState = "pending"	//任务正在审核
	Task_do		TaskState = "doing"		//任务正在进行中
	Task_check	TaskState = "checking"  //任务发布者检查任务完成情况
	Task_other	TaskState = "other"		//其他状况
	Task_finish	TaskState = "finished"	//任务完成
)

//注：原则上，所有的输入数据合法性检查由controller处进行。在models进行的是与模型之间数据关系有关的检查，比如任务的最多接纳人数是否达到上界，等。
type Task struct{
	Id				int			`json:"tid"`
	Userid			int			`json:"userid"`
	Type			string		`json:"type"`//原则上是不接受空格的，表示任务属于某个类型，类型之间互斥
	Description		string		`json:"description"`
	Reward			float32		`json:"reward"`
	Deadline 		string		`json:"deadline"`
	Label			string		`json:"label"`//原则上不接受label带空格，label与label之间使用空格分隔，标签之间不互斥
	State			TaskState			`json:"state" orm:"default(0)"`
	Priority		int32				`json:"priority" orm:"default(0)"`//采用linux优先级策略，越小优先级越高，范围为-255~+255，一般默认为0
	MaxAccept		int32 				`json:"maxAccept" orm:"default(1)"`//任务同时允许的最大接受人数
	HasAccept		int32				`json:"hasAccept" orm:"default(0)"`
	AcceptRelation	[]*AcceptRelation	`json:"acceptRelation" orm:"reverse(many)"`
	ReleaseRelation []*ReleaseRelation	`json:"releaseRelation" orm:"reverse(many)"`
}

/*
功能函数群
*/

//通过任务id获得任务指针
func GetTaskById(taskId int) (*Task,error){
	var tasks []*Task
	o := orm.NewOrm()
	
	if num,err := o.QueryTable("task").Filter("id",taskId).All(&tasks); err != nil || num == 0{
		return nil,fmt.Errorf("task id not found")
	} else if num>1 {
		return nil,fmt.Errorf("task id duplicate")
	}else{
		return tasks[0],nil
	}
}
//根据userId，拿出创建者为同一人的任务
func GetTaskByUserid(userId int) ([]*Task,error){
	var tasks []*Task
	o := orm.NewOrm()
	
	if num,err := o.QueryTable("task").Filter("userid",userId).All(&tasks); err != nil || num == 0{
		return nil,fmt.Errorf("this user don't have any task")
	} else{
		return tasks,nil
	}
}
/*
业务函数群
*/

/*
函数目的：创建任务
调用时机：主要由controllers/task.go调用，实际上可用于任何时刻的添加
需要执行的任务：
	1.向数据库中写入数据

调用成功：直接返回taskId与nil
调用失败：返回-1与err
	可能场景：重复的taskId（如果不是用户指定的，则不会有这种情况）
*/
func AddTask(task *Task) (taskId int,err error){
	o := orm.NewOrm()
	id64,err := o.Insert(task)
	taskId = int(id64)
	if err == nil{
		return taskId,nil
	} else{
		return -1,err
	}
}
/*
函数目的：获取任务
调用时机：任何需要通过taskId获取任务的场景
需要执行的任务：
	1.根据给定的taskId获取task对象

调用成功：直接返回task指针与nil
调用失败：返回nil与错误对象
	可能场景：不存在taskId，或是taskId对应的任务状态为已删除
*/

func GetTask(taskId int) (task *Task,err error){
	o := orm.NewOrm()
	task = &Task{Id:taskId}
	err = o.Read(task)
	if err == nil{
		return
	}else{
		task = nil
		return
	}
}

/*
函数目的：更新任务信息
调用时机：任何需要根据taskId与给定任务对象来更新任务信息的场景
需要执行的任务：
	1.根据给定的taskId与任务对象，更新数据中的对应的任务对象

调用成功：直接返回修改过的任务对象与nil
调用失败：返回nil与错误对象
	可能场景：不存在taskId，或是taskId对应的任务状态为已删除
*/

func UpdateTask(taskId int,tt *Task) (task *Task,err error){
	o := orm.NewOrm()
	task,err = GetTask(taskId)
	if err != nil{
		return nil,err
	}
	task.Type = tt.Type
	task.Description = tt.Description 
	task.Reward = tt.Reward
	task.Deadline = tt.Deadline
	task.Label = tt.Label
	task.State = tt.State
	task.Priority =  tt.Priority
	task.MaxAccept = tt.MaxAccept
	task.AcceptRelation = tt.AcceptRelation
	task.ReleaseRelation = tt.ReleaseRelation
	_,err = o.Update(task)
	if err == nil{
		return
	} else{
		task = nil
		return
	}
}
/*
函数目的：删除任务
调用时机：需要根据taskId删除任务的场景
需要执行的任务：
	1.将状态转为已删除

调用成功：返回nil
调用失败：返回err
	调用失败场景：不存在taskId，或是taskId对应的任务状态为已删除
*/
func DeleteTask(taskId int) error{
	o := orm.NewOrm()
	if _,err := o.Delete(&Task{Id:taskId}); err == nil{
		return nil
	} else{
		return err
	}
}

//下面为功能性函数区域

/*
函数目的：修改任务的状态
调用时机：任何需要修改任务的状态的场景，包括关闭任务和删除任务
需要执行的任务：
	1.修改数据库，修改对应任务的任务状态

调用成功：返回nil
调用失败：返回err
	调用失败场景：不存在taskId，或taskId对应的任务状态已经为已删除

func ChangeState(taskId int, ts TaskState) error{

}

/*
函数目的：用户接受任务
调用时机：从controller那边接到用户希望接受任务的请求
需要执行的任务：
	1.根据给定的userId和taskId，拿到对应的对象
	2.创建新的AcceptRelation，更新到数据库内。（因为是外键关系，内容全部保存在AcceptRelation部分，因此不用更新对应的user和task） 不用更新只是推测，还需要进一步证明
	3.将创建好地AcceptRelation加入到数组之中

调用成功：nil
调用失败：返回error
	调用失败场景：不存在userId或taskId，或者userId已经接受了此任务，或者任务的接纳人数已经达到上限

func AcceptTask(userId,taskId int) error{

}

/*
函数目的：用户取消接受任务
调用时机：从controller那边接到用户希望取消接受任务的请求
需要执行的任务：
	1.根据给定的userId和taskId拿到对应的对象
	2.搜索到AcceptRelelation，并删除
	3.（可能）向取消的用户收取违约金

调用成功：nil
调用失败：返回error
	调用失败场景：不存在userId或taskId，或者userId没有接受这个task，这个AcceptRelation不存在。
		或者存在，但是该任务已经完成

func CancelTask(userId,taskId int) error{

}

/*
函数目的：用户完成任务
调用时机：从controller那边接到用户已经完成了任务
需要执行的任务：
	1.添加releaseRelation
	2.可能要修改其他的用户与该任务的关系，向他们发出通知（也可能是共存任务）
	3.修改该任务的状态
	4.向用户发放报酬

调用成功：nil
调用失败：返回error
	调用失败场景：不存在userId或taskId，或者userId没有接受任务，或者该任务已经完成

func FinishTask(userId,taskId int) error{
	
}
*/