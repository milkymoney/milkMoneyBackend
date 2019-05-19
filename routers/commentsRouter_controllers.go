package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetAllTask",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
            Router: `/:taskId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:taskId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:taskId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "AcceptanceCheckFinishTask",
            Router: `/settleup/:taskid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "AcceptTask",
            Router: `/task/accept/:taskId`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
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
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Query",
            Router: `/query`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
