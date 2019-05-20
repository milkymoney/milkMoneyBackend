package controllers

import (
	"github.com/milkymoney/milkMoneyBackend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
)


type TaskController struct {
	beego.Controller
}

/*
测试后发现必须要有如下条目才能够正常读取
{
  "Type": "string",
  "Description": "string",
  "Reward": 0,
  "Deadline": "string",
  "Label":"test"
}
*/

//因为微信只要收到服务器回复，无论http状态码都会直接调用success回调函数，所以需要在返回值里面加上返回信息以帮助前端确认完成状态。
//同时增加容错
type HttpResponseCode struct{
	Message 	string	`json:"message"`
	Success		bool	`json:"success"`
}




// 拿到所有任务，按照页面进行分解。（目前仅能够返回所有的任务，没有做页面，没有做筛选）
// @Title 查询已发布任务列表
// @Description get task by taskId
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	page		query 	integer	true		"page value,default is 1"
// @Param	myrelease	query 	boolean	false		"check release task"
// @Param	myacceptance	query 	boolean	false		"check accept task"
// @Param	keyword		query 	string	false		"search by labels"
// @Success 200 {[object]} models.Task
// @Failure 403 :taskid is empty
// @router / [get]
func (t *TaskController) GetAllTask() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
	} else{
		tasks,err := models.GetTaskByUserid(user.Id)
		if err != nil{
			t.Data["json"] = err.Error()
		} else{
			t.Data["json"] = tasks
		}
	}
	t.ServeJSON()
}
//创建任务API所需要的返回值
type CreateTaskReturnCode struct{
	HttpResponseCode
	TaskId 		int		`json:"taskId"`
}

// @Title 发布任务
// @Description 用户发送发布任务的请求
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} controllers.CreateTaskReturnCode 
// @Failure 403 body is empty
// @router / [post]
func (t *TaskController) Post() {
	fmt.Println("Create Task")
	var task models.Task
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:err.Error(),Success:false},TaskId:task.Id}
	} else{
		task.Userid = user.Id
		fmt.Println(task)
		tId,err := models.AddTask(&task)
		if err == nil{
			t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:"success",Success:true},TaskId:tId}
		} else{
			t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:err.Error(),Success:false},TaskId:tId}
		}
	}
	t.ServeJSON()
}

// @Title 查询指定id的任务
// @Description get task by taskId
// @Param	taskId		path 	integer	true		"the key"
// @Success 200 {object} models.Task
// @Failure 403 :taskid is empty
// @router /:taskId [get]
func (t *TaskController) Get() {
	tid,err := t.GetInt(":taskId")
	if err==nil {
		taskId := int(tid)
		task, err := models.GetTask(taskId)
		if err != nil {
			t.Data["json"] = err.Error()
		} else {
			t.Data["json"] = task
		}
	} else{
		t.Data["json"] = err.Error()
	}
	t.ServeJSON()
}

// @Title 修改任务信息
// @Description 发布者任务被接取前可以随时修改任务信息，但是在任务被接受后则不能有任何改动
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"The taskId you want to update"
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 :taskid is not int
// @router /:taskId [put]
func (t *TaskController) Put() {
	tid,err := t.GetInt(":taskId")
	if err!=nil {
		taskId := int(tid)

		var task models.Task
		json.Unmarshal(t.Ctx.Input.RequestBody, &task)
		
		_ , err := models.UpdateTask(taskId, &task)
		if err != nil {
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		} else {
			t.Data["json"] = HttpResponseCode{Success:true,Message:"update success"}
		}

	}else{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	}
	t.ServeJSON()
}

// @Title 删除任务
// @Description 考虑到已经发布的任务可能存在因为其他不可抗逆因素而需要取消的情况，任务允许删除。条件是该任务未被接受或者被接受时间不超过十分钟
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"The taskid you want to delete"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 taskid is empty
// @router /:taskId [delete]
func (t *TaskController) Delete() {
	tid,err := t.GetInt(":taskId")
	if err ==nil{
		taskId := int(tid)
		models.DeleteTask(taskId)
		t.Data["json"] = HttpResponseCode{Success:true,Message:"delete success"}
	} else{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	}

	t.ServeJSON()
}
//注：暂时不检查是否达到最大值，只是进行处理
// @Title 接受任务
// @Description user accept the task
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /task/accept/:taskId [post]
func (t *TaskController) AcceptTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	} else{
		tid,err := t.GetInt(":taskId")
		if err != nil{
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		}else{
			taskId := int(tid)
			_,err = models.CreateNewAcRelById(user.Id,taskId,"20190513")//暂时还没有加上时间
			if err == nil{
				t.Data["json"] = HttpResponseCode{Success:true,Message:"accept success"}
			} else{
				t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
			}
		}
	}
	t.ServeJSON()
}

//暂时没做图片证明部分prove

type AcceptorCheckFinishCodeResponse struct{
	Proves			[]string	`json:"proves"`
	models.Task
}

// @Title 接受者查询自己已完成任务的信息
// @Description 接受者完成任务的信息以及证明只有发布者和接受者自身才能看见
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /settleup/:taskid [get]
func (t *TaskController) AcceptorCheckFinishTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	} else{
		tid,err := t.GetInt(":taskId")
		if err != nil{
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		}else{
			taskId := int(tid)
			//拿到了任务id，现在拿到任务，并返回数组
			task,err := models.GetTask(taskId)
			if err != nil{
				t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
			} else{
				images,err := models.GetImagesByUserAndTaskId(user.Id,task.Id)
				if err !=nil{
					t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
				} else{
					//将images数组中的路径全部取出，打包成新的数组
					imageUrl := make([]string,len(images))
					for _,image := range images{
						imageUrl = append(imageUrl,image.ImagePath)
					}
					t.Data["json"] = AcceptorCheckFinishCodeResponse{Task:*task,Proves:imageUrl}
				}
			}
		}
	}
	t.ServeJSON()
}

// @Title 任务接受者结算任务
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Param   graph		body	binary	true		"图片，使用myfile:xxx传输"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /settleup/:taskId [post]
func (t *TaskController) ExecutorSettleupTask(){
	user,err := Auth(&t.Controller)//拿到用户登陆信息
	if err == nil{
		tid,err := t.GetInt(":taskId")
		if err != nil{
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		} else{
			f,h,err := t.GetFile("myfile")
			if err != nil{
				t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
			} else{
				//成功收到图片，将图片的路径保存到本地
				defer f.Close()
				path := "./image/"+h.Filename
				t.SaveToFile("myfile",path)//保存图片到本地

				//顺利拿到关系
				releaseRelation,err := models.GetReleaseRelation(user.Id,int(tid))
				if err != nil{//鬼知道什么错误
					t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
				}else if len(releaseRelation)==0{//没有这个关系
					t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Errorf("no release relation between this user and task.").Error()}
				} else{
					//成功将图片加入到数据库
					err = models.AddImageToSQL(releaseRelation[0].Id,&models.ConfirmImage{ImagePath:h.Filename,ReleaseRelation:releaseRelation[0]})
					if err !=nil{//天晓得什么错误
						t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
					} else{
						t.Data["json"] = HttpResponseCode{Success:true,Message:"success"}
					}
				}
			}
		}

	}else{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	}

	t.ServeJSON()
}


//发布者查询任务完成信息所需要的返回类型
//返回对象时一个数组，单个对象包括用户和其上传的图片数组
type PublisherCheckTaskFinishResponse struct{
	models.User
	Proves 			[]string
}

// @Title 发布者查询任务完成信息
// @Description 任务接受者将任务完成信息上传到服务器上，发布者查看所有接受者的信息以及其上传的任务完成信息
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Success 200 {[object]} controllers.Task
// @Failure 403 {object} controllers.HttpResponseCode
// @router /confirm/:taskId [get]
func (t *TaskController) PublisherCheckTaskFinish(){
	//还没有做用户权限检测
	_,err := Auth(&t.Controller)

	
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	} else{
		tid,err := t.GetInt(":taskId")
		if err != nil{
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		}else{
			taskId := int(tid)
			//拿到了任务id，现在拿到任务，并返回数组
			task,err := models.GetTask(taskId)
			if err != nil{
				t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
			} else{
				//对于给定任务，访问所有的完成情况
				//！！！！！！！！这里有问题
				relations,err := models.GetReleaseRelationByTaskId(task.Id)
				if err != nil{
					t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
				}else{
					//对于所有的完成关系，汇集用户信息和任务信息，制成数组并返回
					ansSet := make([]*PublisherCheckTaskFinishResponse,len(relations))
					for _,relation := range relations{
						//拿到关系对应的用户信息
						aimUser,_ := models.GetUserThroughRelRelation(relation)
						//拿到关系对应的确认图片信息的路径
						images,_ := models.GetImagesByRelationId(relation.Id)
						imageUrl := make([]string,len(images))
						for _,image := range images{
							imageUrl = append(imageUrl,image.ImagePath)
						}
						//添加两者到数组中
						ansSet = append(ansSet,&PublisherCheckTaskFinishResponse{User:*aimUser,Proves:imageUrl})
					}
					//返回数组
					t.Data["json"] = ansSet
				}
				
			}
		}
	}
	t.ServeJSON()
}

// @Title 发布者结算任务
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /confirm/:taskId [post]
func (t *TaskController) PublisherFinishTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	} else{
		tid,err := t.GetInt(":taskId")
		if err != nil{
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		}else{
			taskId := int(tid)
			//拿到了任务id，现在拿到任务，并返回数组
			task,err := models.GetTask(taskId)
			if err != nil{
				t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
			} else{
				//对于给定任务，访问所有的完成情况
				relations,err := models.GetReleaseRelation(user.Id,task.Id)
				if err != nil{
					t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
				}else{
					//需要修改任务的状态
					//因为暂时还没有确定，就不作了
					relations[0] = nil
					t.Data["json"] = HttpResponseCode{Success:true,Message:"haven't done yet"}
				}
				
			}
		}
	}
	t.ServeJSON()
}