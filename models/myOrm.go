package models

import(
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	fmt.Println("Begin to connect to sql")
    orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "wty:97112500@tcp(127.0.0.1:3306)/test?charset=utf8")
	orm.RegisterModel(new(User),new(Task),new(AcceptRelation),new(ReleaseRelation))
	orm.RunSyncdb("default", false, true)
	fmt.Println("Connect over")
}

