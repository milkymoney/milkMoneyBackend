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
	var runmode = beego.AppConfig.String("runmode")
	if runmode == "test"{
		userId,err := u.GetInt("userId")
		if err != nil{
			return nil,fmt.Errorf("Login error")
		} else{
			user,err := models.GetUserById(userId)
			return user,err
		}
	} else{
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


//测试用函数之登陆验证
// @Title login
// @Description Use session to get user's id
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /query [get]
func (u *UserController) Query() {
	fmt.Println("In Query")
	session := u.Ctx.Input.CruSession
	fmt.Println("In controller's function query, get the session")
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

//测试用函数之图片上传
// @Title login
// @Description Use session to get user's id
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /queryImage [post]
func (u *UserController) QueryImage() {
	fmt.Println("In Query")
	session := u.Ctx.Input.CruSession
	if val := session.Get("openid"); val != nil {
		user,err := models.GetUserByOpenId(val.(string))
		if err == nil {
			fmt.Println("user id")
			fmt.Println(user.Id)
			f,h,err := u.GetFile("myfile")
			fmt.Println(f)
			if err != nil{
				fmt.Println("get file err",err)
			} else{
				//成功收到图片
				path := "./image/"+h.Filename
				u.SaveToFile("myfile",path)//保存图片到本地
				fmt.Println("Add path to file:")
				fmt.Println(path)
				models.AddImageToUser(user.Id,path)
				defer f.Close()
			}
			u.Data["json"] = "receive"
		}
	}

	u.ServeJSON()
}

//测试用函数之图片下载
// @Title login
// @Description Use session to get user's id
// @Param	code	query 	string	true		"wx.Login response code"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /download [get]
func (u *UserController) DownloadImage() {
	fmt.Println("In Query")
	session := u.Ctx.Input.CruSession
	if val := session.Get("openid"); val != nil {
		user,err := models.GetUserByOpenId(val.(string))
		if err == nil {
			fmt.Println("user id")
			fmt.Println(user.Id)
			path := models.GetImageFromUser(user.Id)
			u.Ctx.Output.Download(path,"test.png")
		}
	}

}

// @Title 增加积分
// @Description add balance
// @Param	session		header 	string	true
// @Param   money		query	string	true
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /money [post]
func (u *UserController) AddMoney() {
	user,err := Auth(&u.Controller)
	if err != nil{
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	} 

	money,err := u.GetInt("money")
	if err != nil{
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	} 
	user.Balance += money
	_,err = models.UpdateUser(user.Id,user)
	if err != nil{
		u.Data["json"] = err.Error()
		u.ServeJSON()
		return
	} 
	u.ServeJSON()
}