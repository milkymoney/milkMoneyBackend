package models

import (
	"errors"
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init() {

}

type User struct {
	Id      		string				`orm:"pk"`
	Username		string
	Password		string
	Balance			int
	AcceptRelation 	[]*AcceptRelation	`orm:"rel(fk);null;on_delete(do_nothing)"`
	ReleaseRelation []*ReleaseRelation	`orm:"rel(fk);null;on_delete(do_nothing)"`
}


func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	o := orm.NewOrm()
	o.Insert(&u)
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	o := orm.NewOrm()
	u = &User{Id:uid}
	fmt.Println("Ready to get user ",uid)
	err = o.Read(u)
	fmt.Println("Finish to get user ",uid)
	if err == nil{
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	o := orm.NewOrm()
	u,err := GetUser(uid)
	if err == nil{
		if uu.Username != ""{
			u.Username = uu.Username
		}
		if uu.Password != ""{
			u.Password = uu.Password
		}
		if uu.Balance != 0{
			u.Balance = uu.Balance
		}
		o.Update(u)
		return u,nil
	}
	return nil, errors.New("User Not Exist")
}
/*
func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
*/
func DeleteUser(uid string) {
	o := orm.NewOrm()
	u,err := GetUser(uid)
	if err == nil{
		_,err = o.Delete(u)
	}
}
