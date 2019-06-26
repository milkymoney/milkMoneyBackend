package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	fmt.Println("Begin to connect to sql")
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:97112500@tcp(114.115.163.164:30000)/demo?charset=utf8")
	orm.RegisterModel(new(User),new(Task),new(AcceptRelation),new(ReleaseRelation),new(ConfirmImage))
	orm.RunSyncdb("default", false, true)
	test()
}

func test(){
	user,_ := GetUserByName("wty")
	fmt.Println(user)
}

