package controllers

import (
	"apiproject/models"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
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
	uid,err := models.AddUser(user)
	if err == nil{
		userId := strconv.Itoa(uid)
		u.Data["json"] = map[string]string{"uid": userId}
	} else{
		u.Data["json"] = err.Error()
	}
	
	u.ServeJSON()
}

// @Title Get
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

// @Title Update
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

// @Title Delete
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
