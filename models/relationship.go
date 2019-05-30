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
	ConfirmImages	[]*ConfirmImage	`orm:"reverse(many)"`
	User		*User	`orm:"rel(fk)"`
	Task		*Task	`orm:"rel(fk)"`
}
/*
发布者任务关系
*/
type ReleaseRelation struct{
	Id				int
	ReleaseDate		string
	User			*User			`orm:"rel(fk)"`
	Task			*Task			`orm:"rel(fk)"`
}

//完成任务时确认要用到的图片，最多三个
type ConfirmImage struct{
	Id					int
	ImagePath			string //保存的实际上应该是文件名，比如xxx.png，访问的时候目前是可以通过静态文件访问，比如域名/image/xxx.png这样子
	AcceptRelation		*AcceptRelation	`orm:"rel(fk)"`
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
	_,err := o.QueryTable("confirm_image").Filter("release_relation_id",relationId).All(&images)
	return images,err
}
//根据用户id和任务id拿到图片数组
func GetImagesByUserAndTaskId(userId,taskId int) ([]*ConfirmImage,error){
	relations,err := GetReleaseRelation(userId,taskId)
	if err!=nil{
		return nil,err
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
	newRelation := &AcceptRelation{AcceptDate:acceptDate,User:user,Task:task}
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
	newRelation := &ReleaseRelation{ReleaseDate:releaseDate,User:user,Task:task}
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
		return nil,fmt.Errorf("User id or task id not correct.")
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
		return nil,fmt.Errorf("User id or task id not correct.")
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