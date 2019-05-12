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

// @Title GetUser
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		userId,err := strconv.Atoi(uid)
		if err == nil{
			user, err := models.GetUser(userId)
			if err != nil {
				u.Data["json"] = err.Error()
			} else {
				u.Data["json"] = user
			}
		} else{
			u.Data["json"] = err.Error()
		}

	}
	u.ServeJSON()
}

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
	if openid,err := models.Login(code);err==nil {
		//设置session
		u.Data["json"] = openid
		session.Set("openid",openid)
	} else {
		u.Data["json"] = err.Error()
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

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
