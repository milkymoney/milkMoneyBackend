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
            Method: "GetAllTaskPublish",
            Router: `/publisher`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "PublishTask",
            Router: `/publisher`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetPublishTask",
            Router: `/publisher/:taskId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "PutPublishTask",
            Router: `/publisher/:taskId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "DeletePublishTask",
            Router: `/publisher/:taskId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "PublisherConfirmTask",
            Router: `/publisher/confirm/:taskId`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "PublisherCheckTaskFinish",
            Router: `/publisher/confirm/:taskId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetAllTaskAccept",
            Router: `/recipient`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "GetUserAcTask",
            Router: `/recipient/:taskId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "AcceptTask",
            Router: `/recipient/:taskId`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "AcceptorCheckFinishTask",
            Router: `/recipient/settleup/:taskId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:TaskController"],
        beego.ControllerComments{
            Method: "ExecutorSettleupTask",
            Router: `/recipient/settleup/:taskId`,
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
            Method: "AddMoney",
            Router: `/balance`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "DownloadImage",
            Router: `/download`,
            AllowHTTPMethods: []string{"get"},
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
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/milkymoney/milkMoneyBackend/controllers:UserController"],
        beego.ControllerComments{
            Method: "QueryImage",
            Router: `/queryImage`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
