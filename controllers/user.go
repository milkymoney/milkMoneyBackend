package controllers

import (
	"github.com/milkymoney/milkMoneyBackend/models"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
	"fmt"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

/*
	输入当前的beego控制器
	根据session返回session中openid对应的用户指针，或者错误
*/
func Auth(u *beego.Controller) (*models.User,error){
	session := u.Ctx.Input.CruSession
	if val := session.Get("openid"); val != nil {
		user,err := models.GetUserByOpenId(val.(string))
		if err != nil{
			return nil,err
		}else{
			return user,nil
		}
	} else {
		return nil,fmt.Errorf("need login")
	}

}
//待处理，不需要创建用户了，估计就不要了
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	uid,err := models.AddUser(&user)
	if err == nil{
		userId := strconv.Itoa(uid)
		u.Data["json"] = map[string]string{"uid": userId}
	} else{
		u.Data["json"] = err.Error()
	}
	
	u.ServeJSON()
}
//待处理，改了，但是不知道微信登陆用不用的了
// @Title GetUser
// @Description get user by openid(in session)
// @Param	openid		header 	string	true		"user's id from wx"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router / [get]
func (u *UserController) Get() {
	user,err := Auth(&u.Controller)
	if err != nil{
		u.Data["json"] = err.Error()
	} else{
		u.Data["json"] = user
	}
	u.ServeJSON()
}
//待处理，可以留着，但是还没改
// @Title UpdateUser
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		userId,err := strconv.Atoi(uid)
		if err == nil{
			uu, err := models.UpdateUser(userId, &user)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = uu
			}
		} else{
			u.Data["json"] = err.Error()
		}

	}
	u.ServeJSON()
}
//待处理，估计会不要
// @Title DeleteUser
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	userId,err := strconv.Atoi(uid)
	if err ==nil{
		models.DeleteUser(userId)
		u.Data["json"] = "delete success!"
	} else{
		u.Data["json"] = err.Error()
	}

	u.ServeJSON()
}
//待处理，未与微信同步
// @Title Login
// @Description Logs user into the system
// @Param	code		query 	string	true		"the code from wx.login()"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	fmt.Println(u.Ctx.Input.CruSession)
	session := u.Ctx.Input.CruSession
	code := u.GetString("code")
	fmt.Println("Code" + code)
	if openid,err := models.Login(code);err==nil {
		//设置session
		u.Data["json"] = openid
		session.Set("openid",openid)
	} else {
		u.Data["json"] = err.Error()
	}
	fmt.Println("Set session over")
	u.ServeJSON()
}

//待处理，可能不要
// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

//测试用函数
// @Title login
// @Description Use session to get user's id
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /query [get]
func (u *UserController) Query() {
	fmt.Println("In Query")
	session := u.Ctx.Input.CruSession
	fmt.Println(session)
	if val := session.Get("openid"); val != nil {
		user,err := models.GetUserByOpenId(val.(string))
		fmt.Println(user)
		if err !=nil{
			u.Data["json"] = "openid error"
		} else{
			u.Data["json"] = user.Id
		}
	} else {
		u.Data["json"] ="need login"
	}
	u.ServeJSON()
}
