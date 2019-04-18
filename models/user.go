package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {

}

type User struct {
	Id      		int
	Username		string
	Password		string
	Balance			int
	AcceptRelation	[]*AcceptRelation	`orm:"reverse(many)"`
	ReleaseRelation []*ReleaseRelation	`orm:"reverse(many)"`
}

/*
函数目的：创建用户
调用时机：添加用户的时候
需要执行的任务：
	1.向数据库中写入数据

调用成功：直接返回userId与nil
调用失败：返回-1与err
	可能场景：重复的userId（如果不是用户指定的，则不会有这种情况）
*/
func AddUser(user User) (userId int,err error) {
	o := orm.NewOrm()
	id64,err := o.Insert(&user)
	userId = int(id64)
	if err == nil{
		return userId,nil
	} else{
		return -1,err
	}
}

/*
函数目的：获取用户
调用时机：根据userId获取用户信息
需要执行的任务：
	1.向数据库中读取信息

调用成功：直接返回User指针与nil
调用失败：返回nil与err
	可能场景：不存在userId
*/

func GetUser(userId int) (u *User, err error) {
	o := orm.NewOrm()
	u = &User{Id:userId}
	err = o.Read(u)
	if err == nil{
		return
	}else{
		u = nil
		return
	}
}
/*
函数目的：更新用户信息
调用时机：任何需要根据userId与给定用户对象来更新用户信息的场景
需要执行的任务：
	1.根据给定的userId与用户对象更新指定用户的信息。

具体实现：
	更新除了Id与AcceptRelation和ReleaseRelation数组外其他所有的信息

调用成功：直接返回修改过的用户对象与nil
调用失败：返回nil与错误对象
	可能场景：不存在对应的用户
*/
func UpdateUser(userId int, uu *User) (user *User, err error) {
	o := orm.NewOrm()
	user,err = GetUser(userId)
	if err != nil{
		return nil,err
	}
	user.Username = uu.Username
	user.Password = uu.Password
	user.Balance = uu.Balance
	_,err = o.Update(user)
	if err == nil{
		return
	} else{
		user = nil
		return
	}
}

/*
函数目的：删除用户
调用时机：需要根据userId删除用户的场景
需要执行的任务：
	1.将用户从数据库中删除

调用成功：返回nil
调用失败：返回err
	调用失败场景：不存在userId
*/
func DeleteUser(userId int) error {
	o := orm.NewOrm()
	if _,err := o.Delete(&User{Id:userId}); err == nil{
		return nil
	} else{
		return err
	}
}
