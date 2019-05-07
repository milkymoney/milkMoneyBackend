package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	fmt.Println("Begin to connect to sql")
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wu97112500@tcp(129.204.7.185:30000:30000)/demo?charset=utf8")
	orm.RegisterModel(new(User),new(Task),new(AcceptRelation),new(ReleaseRelation))
	orm.RunSyncdb("default", false, true)
	test()
}

func test(){
	acRelation,err := CreateNewAcRelById(1,4,"2017-08-08")
	fmt.Println(acRelation)
	fmt.Println(err)
}

