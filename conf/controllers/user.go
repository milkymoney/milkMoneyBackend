package controllers

import (
	"github.com/milkymoney/milkMoneyBackend/models"
	"encoding/json"
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

// @Title 查询用户信息
// @Description get user by openid(in session)
// @Param	session		header 	string	true		"user's session ,get from login"
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
// @Title 修改用户个人信息
// @Description update the user
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	originUser,err := Auth(&u.Controller)
	if err == nil {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		userId := originUser.Id
		if err == nil{
			_, err := models.UpdateUser(userId, &user)
			if err != nil {
				u.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
			} else {
				u.Data["json"] = HttpResponseCode{Success:true,Message:"update success"}
			}
		} else{
			u.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		}

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
	fmt.Println("In controller function login, it's session:")
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


//测试用函数
// @Title login
// @Description Use session to get user's id
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /query [post]
func (u *UserController) Query() {
	fmt.Println("In Query")
	session := u.Ctx.Input.CruSession
	var user models.User
	if val := session.Get("openid"); val != nil {
		user,err := models.GetUserByOpenId(val.(string))
		if err == nil {
			fmt.Println(user)
		}
	}
	fmt.Println("user check")
	fmt.Println(user)
	f,h,err := u.GetFile("myfile")
	fmt.Println(f)
	if err != nil{
		fmt.Println("get file err",err)
	} else{
		u.SaveToFile("myfile","./image/"+h.Filename)
		defer f.Close()
	}
	u.Data["json"] = "receive"
	u.ServeJSON()
}

