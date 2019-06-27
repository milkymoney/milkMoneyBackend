package models

import(
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init(){

}



type TaskState int

const(
	Task_ac_pend 	TaskState = 0	//任务正在审核
	Task_ac_do		TaskState = 1		//任务正在进行中
	Task_ac_check	TaskState = 2  //任务发布者检查任务完成情况
	Task_ac_finish	TaskState = 3		//其他状况
	Task_rel_pend	TaskState = 4	//任务完成
	Task_rel_do		TaskState = 5
	Task_rel_finish	TaskState = 6
)

const(
	Check_uncheck	string = "unchecked"
	Check_pass		string = "passed"
	Check_unpass	string = "unpassed"
)

/*接受关系为用户到task的多对多关系，一个用户可以接受多个任务，一个任务也可以被多个用户所接受
但是实现上为一对一关系，即用户id与任务id绑定，一起组成acId
原则上一个用户不能够接受*/
type AcceptRelation struct {
	Id				int
	AcceptDate		string
	ConfirmImages	[]*ConfirmImage	`orm:"reverse(many)"`
	AcTaskState		TaskState
	CheckState		string
	User			*User	`orm:"rel(fk)"`
	Task			*Task	`orm:"rel(fk)"`
}
/*
发布者任务关系
*/
type ReleaseRelation struct{
	Id				int
	ReleaseDate		string
	RelTaskState	TaskState
	User			*User			`orm:"rel(fk)"`
	Task			*Task			`orm:"rel(fk)"`
}

//完成任务时确认要用到的图片，最多三个
type ConfirmImage struct{
	Id					int
	ImagePath			string //保存的实际上应该是文件名，比如xxx.png，访问的时候目前是可以通过静态文件访问，比如域名/image/xxx.png这样子
	AcceptRelation		*AcceptRelation	`orm:"rel(fk)"`
}

//拿到所有的已发布任务

func GetAllPublishTask() ([]*Task,error){
	var reRelations []*ReleaseRelation
	o := orm.NewOrm()
	_,err := o.QueryTable("release_relation").All(&reRelations)
	if err != nil{
		return nil,err
	}
	//将拿到的关系中的id用来
	var tasks []*Task
	for _,relation := range(reRelations){
		o.Read(relation.Task)
		tasks = append(tasks,relation.Task)
	}
	return tasks,nil
}

//通过用户id拿去用户发布和接收的所有任务
func GetAllTaskByUserid(userId int) ([]*Task,error){
	var acRelations []*AcceptRelation
	var reRelations []*ReleaseRelation
	o := orm.NewOrm()
	_,err := o.QueryTable("accept_relation").Filter("user_id",userId).All(&acRelations)
	if err != nil{
		return nil,err
	}
	_,err = o.QueryTable("release_relation").Filter("user_id",userId).All(&reRelations)
	if err != nil{
		return nil,err
	}
	//将拿到的关系中的id用来
	var tasks []*Task
	for _,relation := range(acRelations){
		o.Read(relation.Task)
		tasks = append(tasks,relation.Task)
	}
	return tasks,nil
}

/*业务函数群*/

//向图片数组中添加一张图片。如果已经有n张图片，删除最早加入的一张（id最小的一张）。
//失败场景：没有找到对应的关系
func AddImageToSQL(relationId int, image *ConfirmImage) error{
	o := orm.NewOrm()
	relation := &AcceptRelation{Id:relationId}
	err := o.Read(relation)
	if err != nil{
		return err
	} 

	//拿到关系，进行加入操作
	image.AcceptRelation = relation
	maxImageNum := 3
	//检查数量，如果超过则需要删除最早加入的一张
	var images []*ConfirmImage
	num,err := o.QueryTable("confirm_image").Filter("accept_relation_id",relationId).All(&images)
	if err != nil{
		return err
	}

	if int(num) == maxImageNum && maxImageNum > 0{//需要执行删除操作
		minId := images[0].Id
		for i := 1; i <= maxImageNum-1; i++{
			if images[i].Id < minId{
				minId = images[i].Id
			}
		}

		//删除最小的Id最对应的image
		_,_ = o.Delete(&ConfirmImage{Id:minId})
	}
	_,err= o.Insert(image)
	if err !=nil{
		return err
	}

	return nil
}

//根据完成关系的id拿到图片数组
func GetImagesByRelationId(relationId int) ([]*ConfirmImage,error){
	o := orm.NewOrm()
	var images []*ConfirmImage
	_,err := o.QueryTable("confirm_image").Filter("accept_relation_id",relationId).All(&images)
	return images,err
}
//根据用户id和任务id拿到图片数组
func GetImagesByUserAndTaskId(userId,taskId int) ([]*ConfirmImage,error){
	relations,err := GetAcceptRelation(userId,taskId)
	if err!=nil{
		return nil,err
	} else if len(relations)==0{
		return nil,fmt.Errorf("该用户与任务之间不存在接受关系")
	}
	images,err := GetImagesByRelationId(relations[0].Id)
	return images,err
}

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
	newRelation := &AcceptRelation{AcceptDate:acceptDate,AcTaskState:Task_ac_do,User:user,Task:task}
	acId,err := CreateAcceptRelation(newRelation)
	if err != nil{
		return nil,err
	}
	newRelation.Id = acId
	return newRelation,nil
}

//根据用户id和任务id创建新的ReleaseRelation，并加入到数据库中。
func CreateNewReRelById(userId,taskId int,releaseDate string) (*ReleaseRelation,error){
	user,err := GetUserById(userId)
	if err != nil{
		return nil,err
	}
	task,err := GetTaskById(taskId)
	if err != nil{
		return nil,err
	}
	newRelation := &ReleaseRelation{ReleaseDate:releaseDate,RelTaskState:Task_rel_do,User:user,Task:task}
	reId,err := CreateReleaseRelation(newRelation)
	if err != nil{
		return nil,err
	}
	newRelation.Id = reId
	return newRelation,nil
}

//本方法没有经过测试

//根据终止关系拿到其用户信息
//因为关系创建是由外键维持的，应该不会失败
func GetUserThroughRelRelation(relation *ReleaseRelation)	(*User,error){
	o := orm.NewOrm()
	//根据relation拿到对应的User
	if relation.User != nil{
		o.Read(relation.User)
	}
	return relation.User,nil
}
func GetUserThroughAcRelation(relation *AcceptRelation)	(*User,error){
	o := orm.NewOrm()
	//根据relation拿到对应的User
	if relation.User != nil{
		o.Read(relation.User)
	}
	return relation.User,nil
}

//拿去任务状态
func GetAcTaskStateThroughTask(user *User,task *Task) (TaskState,error){
	o := orm.NewOrm()
	var relation AcceptRelation
	err := o.QueryTable("accept_relation").Filter("task_id",task.Id).Filter("user_id",user.Id).One(&relation)
	if err != nil{
		return -1,err
	}else{
		return relation.AcTaskState,nil
	}
}
func GetReTaskStateThroughTask(task *Task) (TaskState,error){
	o := orm.NewOrm()
	var relation ReleaseRelation
	err := o.QueryTable("release_relation").Filter("task_id",task.Id).One(&relation)
	if err != nil{
		return -1,err
	}else{
		return relation.RelTaskState,nil
	}
}

//功能函数，通过User拿到所有接受地任务
//未经过测试
func GetAcTaskByUserid(userId int) ([]*Task,error){
	//确认userId没有问题
	_,err := GetUser(userId)
	if err != nil{
		return nil,err
	}
	var relations []*AcceptRelation
	o := orm.NewOrm()
	if num,err := o.QueryTable("accept_relation").Filter("user_id",userId).All(&relations); err != nil || num == 0{
		return nil,fmt.Errorf("用户没有接受过任务，或者存在其他问题。")
	} else{
		//根据拿到的relations读取tasks
		var tasks []*Task
		for i := 0 ; i < len(relations) ;i++{
			o.Read(relations[i].Task)
			tasks = append(tasks,relations[i].Task)
		}
		return tasks,nil
	}

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
	relation.CheckState = Check_uncheck
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
		return nil,fmt.Errorf("用户id与任务id无法找到接受关系")
	} else{
		return relations,nil
	}

}

func GetAcceptRelationByTaskId(taskId int) (relation []*AcceptRelation,err error){
	var relations []*AcceptRelation
	o := orm.NewOrm()

	if _,err := o.QueryTable("accept_relation").Filter("task_id",taskId).All(&relations); err != nil{
		return nil,fmt.Errorf("用户id与任务id无法找到接受关系")
	} else{
		return relations,nil
	}

}

func UpdateAcceptRelation(relation *AcceptRelation) (*AcceptRelation,error){
	o := orm.NewOrm()
	_,err := o.Update(relation)
	return relation,err
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
*/
func CreateReleaseRelation(relation *ReleaseRelation) (int,error){
	o := orm.NewOrm()
	id64,err := o.Insert(relation)
	id := int(id64)
	if err != nil{
		return 0,err
	}
	return id,nil
}

func UpdateReleaseRelation(relation *ReleaseRelation) (*ReleaseRelation,error){
	o := orm.NewOrm()
	_,err := o.Update(relation)
	return relation,err
}
/*
函数目的：拿到ReleaseRelation
调用时机：需要使用userId和taskId拿到relation
需要执行的任务：
	1.从数据库中查询并返回对象

调用成功：返回这个对象,nil
调用失败：nil,err对象
	调用失败场景：查询不到对应的对象
*/
func GetReleaseRelation(userId,taskId int) (relation []*ReleaseRelation,err error){
	var relations []*ReleaseRelation
	o := orm.NewOrm()

	if _,err := o.QueryTable("release_relation").Filter("user_id",userId).Filter("task_id",taskId).All(&relations); err != nil{
		return nil,fmt.Errorf("用户id与任务id无法找到发布关系.")
	} else{
		return relations,nil
	}

}

/*
拿到所有与指定task相关联的ReleaseRelation
*/
func GetReleaseRelationByTaskId(taskId int) (relation []*ReleaseRelation,err error){
	var relations []*ReleaseRelation
	o := orm.NewOrm()

	if _,err := o.QueryTable("release_relation").Filter("task_id",taskId).All(&relations); err != nil{
		return nil,fmt.Errorf("用户id与任务id无法找到发布关系.")
	} else{
		return relations,nil
	}
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
