package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:taskid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:taskid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:taskid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
