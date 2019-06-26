package controllers

import (
	"github.com/milkymoney/milkMoneyBackend/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"strings"
	"strconv"
	"time"
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


//用于查询任务的返回结果
type GetResponse struct{
	*models.Task
	models.TaskState	`json:"state"`
}

//没有测试
//功能函数：给定一个model.Task指针和一个label，要求Task的label满足给定label的所有条目（超集）
func HasLabel(task *models.Task,label string) bool{
	requireString := strings.Split(label," ")
	aimString := strings.Split(task.Label," ")
	for _,s := range requireString{
		if s == ""{
			continue
		}
		pass := false
		//对于给定的label，只要目标数组中存在，就通过。不然，就是不通过.
		for _,aim := range aimString{
			if s == aim{
				pass = true
				break
			}
		}
		if !pass{
			return false
		}
	}
	return true
}

//
func GetUnfinishedTask(tasks []*models.Task) ([]*models.Task,error){

	for i:=0;i<len(tasks);i++{
		reState,err := models.GetReTaskStateThroughTask(tasks[i])
		if err == nil{
			if  reState==models.Task_rel_pend || reState==models.Task_rel_finish{
				//删除不能够显示的任务
				fmt.Println("delete task")
				fmt.Println(reState)
				tasks[i] = tasks[len(tasks)-1]
				tasks = tasks[:len(tasks)-1]
				i--
			}
		}

	}
	return tasks,nil
}


// 拿到所有任务（未结束的已发布任务）
// @Title 查询当前所有未结束的任务
// @Description get task by taskId
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	page		query 	integer	true		"page value,default is 1"
// @Param	keyword		query 	string	false		"search by labels"
// @Success 200 {[object]} models.Task
// @Failure 403 :taskId is empty
// @router / [get]
func (t *TaskController) GetAllTask() {

	tasks,err := models.GetAllPublishTask()
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}
	//筛选，去除已经完成的任务
	tasks,err = GetUnfinishedTask(tasks)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}

	labels := t.GetString("keyword")
	//依据label进行筛选
	if labels != ""{
		num := len(tasks)
		for i:=0;i<num;i++{
			//如果任务不包含标签
			if(!HasLabel(tasks[i],labels)){
				//将该元素与最后一个交换，删除最后一个元素
				tasks[i] = tasks[num-1]
				tasks = tasks[:num-1]
				num--
				i--
			}
		}
	}
	
	//根据页数进行返回
	elementNum := 10
	page := t.GetString("page")
	if page == ""{
		t.Data["json"] = tasks
	} else{
		pageNumber,err := strconv.Atoi(page)
		beginNum := pageNumber*elementNum
		endNum := (pageNumber+1)*elementNum
		if beginNum > len(tasks){
			t.Data["json"] = []*models.Task{}
		} else{
			if endNum > len(tasks){
				endNum = len(tasks)
			}
				
			if err != nil{
				t.Data["json"] = err.Error()
			} else{
				t.Data["json"] = tasks[beginNum:endNum]
			}
		}
	}

	
	t.ServeJSON()
}

// 以下两个为批量查找任务

// @Title 查询自己发布任务列表
// @Description get task by taskId
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	page		query 	integer	true		"page value,default is 1"
// @Success 200 {[object]} models.Task
// @Failure 403 :taskId is empty
// @router /publisher [get]
func (t *TaskController) GetAllTaskPublish() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	tasks,err := models.GetTaskByUserid(user.Id)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

		//根据页数进行返回
	elementNum := 10
	page := t.GetString("page")
	if page == ""{
		var result []GetResponse
		for i:=0 ;i<len(tasks);i++{
			state,_ := models.GetReTaskStateThroughTask(tasks[i])
			result = append(result,GetResponse{Task:tasks[i],TaskState:state})
		}
		t.Data["json"] = result
	} else{
		pageNumber,err := strconv.Atoi(page)
		beginNum := pageNumber*elementNum
		endNum := (pageNumber+1)*elementNum
		if beginNum > len(tasks){
			t.Data["json"] = []*models.Task{}
		} else{
			if endNum > len(tasks){
				endNum = len(tasks)
			}
				
			if err != nil{
				t.Data["json"] = err
			} else{
				var result []GetResponse
				for i:=beginNum ; i<endNum; i++{
					state,_ := models.GetReTaskStateThroughTask(tasks[i])
					result = append(result,GetResponse{Task:tasks[i],TaskState:state})
				}
				t.Data["json"] = result
			}
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
// @router /publisher [post]
func (t *TaskController) PublishTask() {
	var task models.Task
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	task.Id = -1
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:err.Error(),Success:false},TaskId:task.Id}
		t.ServeJSON()
		return
	}
	fmt.Println("Get task",task)
	task.Userid = user.Id
	//检查是否存在足够的资金
	allPay := task.Reward * task.MaxAccept
	if allPay == 0{
		t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:fmt.Sprintf("reward and maxaccept can't be 0"),Success:false},TaskId:task.Id}
		t.ServeJSON()
		return
	} else if user.Balance < allPay{
		t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:fmt.Sprintf("User don't have enough money"),Success:false},TaskId:task.Id}
		t.ServeJSON()
		return	
	}
	user.Balance -= allPay
	_,err = models.UpdateUser(user.Id,user)

	//发布任务的时候保证一些属性不会发生改变
	task.HasAccept = 0
	task.FinishNum = 0
	task.Id = 0
	tId,err := models.AddTask(&task)
	//发布任务，创建发布关系
	_,err = models.CreateNewReRelById(user.Id,tId,time.Now().Format("2006-01-02 15:04:05"))
	if err == nil{
		t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:"success",Success:true},TaskId:tId}
	} else{
		t.Data["json"] = CreateTaskReturnCode{HttpResponseCode:HttpResponseCode{Message:err.Error(),Success:false},TaskId:tId}
	}
	
	t.ServeJSON()
}



// @Title 查询自己已经发布的指定id的任务
// @Description get task by taskId
// @Param	taskId		path 	integer	true		"the key"
// @Success 200 {object} models.GetResponse
// @Failure 403 :taskId is empty
// @router /publisher/:taskId [get]
func (t *TaskController) GetPublishTask() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}
	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err!=nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	task, err := models.GetTask(taskId)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}

	//检验user与任务发布者的关系
	relations,err := models.GetReleaseRelation(user.Id,taskId)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} else if len(relations) == 0{
		t.Data["json"] = fmt.Sprintf("Release Relation of user %d and task %d not exist",user.Id,task.Id)
		t.ServeJSON()
		return	
	}else if len(relations)!=1{
		t.Data["json"] = fmt.Sprintf("Release Relation of user %d and task %d are multiple",user.Id,task.Id)
		t.ServeJSON()
		return	
	}
	
	state,err := models.GetReTaskStateThroughTask(task)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}

	t.Data["json"] = GetResponse{Task:task,TaskState:state}

	t.ServeJSON()
}

// @Title 修改任务信息
// @Description 发布者任务被接取前可以随时修改任务信息，但是在任务被接受后则不能有任何改动
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"The taskId you want to update"
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 :taskId is not int
// @router /publisher/:taskId [put]
func (t *TaskController) PutPublishTask() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}

	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err!=nil {
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}

	var task models.Task
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	
	originTask,err := models.GetTaskById(taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 

	if originTask.Userid != user.Id{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("task not publish by that user.")}
		t.ServeJSON()
		return
	}

	//最终更新任务信息
	_ , err = models.UpdateTask(taskId, &task)
	if err != nil {
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	} else {
		t.Data["json"] = HttpResponseCode{Success:true,Message:"update success"}
	}

	
	t.ServeJSON()
}

// @Title 删除任务
// @Description 考虑到已经发布的任务可能存在因为其他不可抗逆因素而需要取消的情况，任务允许删除。条件是该任务未被接受或者被接受时间不超过十分钟
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"The taskId you want to delete"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 taskId is empty
// @router /publisher/:taskId [delete]
func (t *TaskController) DeletePublishTask() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err !=nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}

	originTask,err := models.GetTaskById(taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	} else if originTask.Userid != user.Id{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("task not publish by that user.")}
	}else{
		models.DeleteTask(taskId)
		t.Data["json"] = HttpResponseCode{Success:true,Message:"delete success"}
	}
	
	t.ServeJSON()
}

//发布者查询任务完成信息所需要的返回类型
//返回对象时一个数组，单个对象包括用户和其上传的图片数组
type PublisherCheckTaskFinishResponse struct{
	models.User
	Proves 			[]string	`json:"proves"`
	CheckState		string		`json:"checkState"`
}

// @Title 发布者查询任务完成信息
// @Description 任务接受者将任务完成信息上传到服务器上，发布者查看所有接受者的信息以及其上传的任务完成信息
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Success 200 {[object]} controllers.Task
// @Failure 403 {object} controllers.HttpResponseCode
// @router /publisher/confirm/:taskId [get]
func (t *TaskController) PublisherCheckTaskFinish(){
	fmt.Println("In publisher check task.")
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}
	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}

	//拿到了任务id，现在拿到任务，并返回数组
	task,err := models.GetTask(taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} else if task.Userid != user.Id{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("Auth fail")}
		t.ServeJSON()
		return
	}
	//对于给定任务，访问所有的完成情况
	relations,err := models.GetAcceptRelationByTaskId(task.Id)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}
	//对于所有的完成关系，汇集用户信息和任务信息，制成数组并返回
	var ansSet []*PublisherCheckTaskFinishResponse
	for _,relation := range relations{
		//如果接受任务已经完成，则不再出现
		if relation.AcTaskState != models.Task_ac_check{
			continue
		}
		//拿到关系对应的用户信息
		aimUser,_ := models.GetUserThroughAcRelation(relation)
		//拿到关系对应的确认图片信息的路径
		images,_ := models.GetImagesByRelationId(relation.Id)
		var imageUrl []string
		for _,image := range images{
			imageUrl = append(imageUrl,image.ImagePath)
		}
		//添加两者到数组中
		ansSet = append(ansSet,&PublisherCheckTaskFinishResponse{User:*aimUser,Proves:imageUrl,CheckState:relation.CheckState})
	}
	//返回数组
	fmt.Println(ansSet)
	t.Data["json"] = ansSet
	t.ServeJSON()
}

//发布者结算任务专用对象
type PublisherConfirmTaskData struct{
	CheckState	string	`json:"checkState"`
	Users 		[]int	`json:"users"`
}

// @Title 发布者结算任务
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Param	body		body 	PublisherConfirmTaskData	true		"main data is the user"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /publisher/confirm/:taskId [post]
func (t *TaskController) PublisherConfirmTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 
	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}
		//拿到了任务id，现在拿到任务，并返回数组
	task,err := models.GetTask(taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 
	//拿到发布关系，检查用户信息
	releaseRelations,err := models.GetReleaseRelation(user.Id,task.Id)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} else if len(releaseRelations) ==0{
		t.Data["json"] = HttpResponseCode{Success:false,Message:"you are not the publisher of task"}
		t.ServeJSON()
		return
	}else if len(releaseRelations) > 1{
		t.Data["json"] = HttpResponseCode{Success:false,Message:"multiple publish"}
		t.ServeJSON()
		return
	}

	//拿到同意的user列表
	var data PublisherConfirmTaskData
	json.Unmarshal(t.Ctx.Input.RequestBody, &data)
	if len(data.Users) == 0{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("Strange, you don't confirm any user.")}
		t.ServeJSON()
		return
	} else if data.CheckState != models.Check_pass && data.CheckState != models.Check_unpass{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("CheckState must be passed or unpassed.")}
		t.ServeJSON()
		return
	}

	//如果pass，则发钱。如果hasAccept达到MaxAccept，则任务发布状态结束。
	//如果unpass，则hasAccept减少
	//状态都要反应在acceptRelation上
	for _,userId := range data.Users{
		acRelation,err := models.GetAcceptRelation(userId,task.Id)
		if err == nil{
			//如果关系正在被检查而且还没有被检查到
			if acRelation[0].AcTaskState == models.Task_ac_check && acRelation[0].CheckState == models.Check_uncheck{
				acRelation[0].AcTaskState = models.Task_ac_finish
				acRelation[0].CheckState = data.CheckState
				_,_ = models.UpdateAcceptRelation(acRelation[0])

				//通过
				if data.CheckState == models.Check_pass{
					//给人发钱
					acceptance,_ := models.GetUser(userId)
					acceptance.Balance += task.Reward
					_,_ = models.UpdateUser(userId,acceptance)
					//任务完成人数上升
					task.FinishNum += 1
					_,_ = models.UpdateTask(task.Id,task)
					//检查任务完成人数，如果与最大人数相同，则发布者任务结束
					if task.FinishNum == task.MaxAccept{
						releaseRelations[0].RelTaskState = models.Task_rel_finish
						_,_ = models.UpdateReleaseRelation(releaseRelations[0])
					}
					
				}else if data.CheckState == models.Check_unpass{
					//hasAccept减少
					task.HasAccept -= 1
					_,_ = models.UpdateTask(task.Id,task)
				}


			}

		}
	}

	t.Data["json"] = HttpResponseCode{Success:true,Message:"string"}
	t.ServeJSON()
}

// @Title 查询自己已接受任务列表
// @Description get task by taskId
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	page		query 	integer	true		"page value,default is 1"
// @Param	keyword		query 	string	false		"search by labels"
// @Success 200 {[object]} models.Task
// @Failure 403 :taskId is empty
// @router /recipient [get]
func (t *TaskController) GetAllTaskAccept() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	}

	tasks,err := models.GetAcTaskByUserid(user.Id)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	//根据页数进行返回
	labels := t.GetString("keyword")
	//依据label进行筛选
	if labels != ""{
		num := len(tasks)
		for i:=0;i<num;i++{
			//如果任务不包含标签
			if(!HasLabel(tasks[i],labels)){
				//将该元素与最后一个交换，删除最后一个元素
				tasks[i] = tasks[num-1]
				tasks = tasks[:num-1]
				num--
				i--
			}
		}
	}

	elementNum := 10
	page := t.GetString("page")
	if page == ""{
		var result []GetResponse
		for i:=0 ;i<len(tasks);i++{
			state,_ := models.GetAcTaskStateThroughTask(user,tasks[i])
			result = append(result,GetResponse{Task:tasks[i],TaskState:state})
		}
		t.Data["json"] = result
		t.ServeJSON()
		return
	} 
	
	pageNumber,err := strconv.Atoi(page)
	beginNum := pageNumber*elementNum
	endNum := (pageNumber+1)*elementNum
	if beginNum > len(tasks){
		t.Data["json"] = []*models.Task{}
	} else{
		if endNum > len(tasks){
			endNum = len(tasks)
		}
			
		if err != nil{
			t.Data["json"] = err
		} else{
			var result []GetResponse
			for i:=beginNum ; i<endNum; i++{
				state,_ := models.GetAcTaskStateThroughTask(user,tasks[i])
				result = append(result,GetResponse{Task:tasks[i],TaskState:state})
			}
			t.Data["json"] = result
		}
	}
	
	t.ServeJSON()
}



// @Title 查询指定id的任务
// @Description get task by taskId
// @Param	taskId		path 	integer	true		"the key"
// @Success 200 {object} models.Task
// @Failure 403 :taskId is empty
// @router /recipient/:taskId [get]
func (t *TaskController) GetUserAcTask() {
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err!=nil{
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	task, err := models.GetTask(taskId)
	if err != nil {
		t.Data["json"] = err.Error()
		t.ServeJSON()
		return
	} 

	acRelations,err := models.GetAcceptRelation(user.Id,taskId)
	if err != nil{
		t.Data["json"] = err.Error()
	} else if len(acRelations)!=1{
		t.Data["json"] = fmt.Sprintf("task %d and user %d's accept relation not correct",taskId,user.Id)
	}else{
		state,err := models.GetAcTaskStateThroughTask(user,task)
		if err != nil{
			t.Data["json"] = err.Error()
			t.ServeJSON()
			return
		}
		t.Data["json"] = GetResponse{Task:task,TaskState:state}
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
// @router /recipient/:taskId [post]
func (t *TaskController) AcceptTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 
	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}

	task,err := models.GetTaskById(taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}
	//已经拿到了用户id和任务id，检查是否有重复情况
	acceptRelations,err := models.GetAcceptRelation(user.Id,taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} else if len(acceptRelations)>0{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("user %d already accept task %d",user.Id,task.Id)}
		t.ServeJSON()
		return
	}

	if task.MaxAccept > task.HasAccept{
		_,err = models.CreateNewAcRelById(user.Id,taskId,time.Now().Format("2006-01-02 15:04:05"))//暂时还没有加上时间
		if err == nil{
			t.Data["json"] = HttpResponseCode{Success:true,Message:"accept success"}
			task.HasAccept += 1
			_,_ = models.UpdateTask(task.Id,task)
		} else{
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		}
	} else{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("Enough people accept the task.")}
	}
	
	t.ServeJSON()
}

//暂时没做图片证明部分prove

type AcceptorCheckFinishCodeResponse struct{
	Proves			[]string	`json:"proveURL"`
	models.Task
}

// @Title 接受者查询自己已完成任务的信息
// @Description 接受者完成任务的信息以及证明只有发布者和接受者自身才能看见
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /recipient/settleup/:taskId [get]
func (t *TaskController) AcceptorCheckFinishTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 

	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}
	
	//拿到了任务id，现在拿到任务，并返回数组
	task,err := models.GetTask(taskId)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}

	images,err := models.GetImagesByUserAndTaskId(user.Id,task.Id)
	if err !=nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}
	//将images数组中的路径全部取出，打包成新的数组
	var imageUrl []string
	for _,image := range images{
		imageUrl = append(imageUrl,image.ImagePath)
	}
	t.Data["json"] = AcceptorCheckFinishCodeResponse{Task:*task,Proves:imageUrl}
	t.ServeJSON()
}

// @Title 任务接受者结算任务
// @Param	session		header 	string	true		"user's session ,get from login"
// @Param	taskId		path 	integer	true		"任务id"
// @Param   graph		body	binary	true		"图片，使用myfile:xxx传输"
// @Success 200 {object} controllers.HttpResponseCode
// @Failure 403 {object} controllers.HttpResponseCode
// @router /recipient/settleup/:taskId [post]
func (t *TaskController) ExecutorSettleupTask(){
	user,err := Auth(&t.Controller)//拿到用户登陆信息
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	}

	tid := t.GetString(":taskId")
	taskId,err := strconv.Atoi(tid)
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 

	f,h,err := t.GetFile("myfile")
	if err != nil{
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		t.ServeJSON()
		return
	} 

	//成功收到图片，将图片的路径保存到本地
	defer f.Close()
	path := "./image/"+h.Filename
	t.SaveToFile("myfile",path)//保存图片到本地
	//顺利拿到关系
	acceptRelation,err := models.GetAcceptRelation(user.Id,taskId)
	if err != nil{//鬼知道什么错误
		t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
	}else if len(acceptRelation)==0{//没有这个关系，说明没有权限或者任务id错误
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("no release relation between this user and task.")}
	} else if len(acceptRelation)>1{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("task %d and user %d's accept relationship more than 1.",taskId,user.Id)}
	}else if acceptRelation[0].AcTaskState == models.Task_ac_finish{
		t.Data["json"] = HttpResponseCode{Success:false,Message:fmt.Sprintf("task %d accepted by user %d has already finish with status %s.",taskId,user.Id,acceptRelation[0].CheckState)}
	}else{
		//成功将图片加入到数据库
		
		err = models.AddImageToSQL(acceptRelation[0].Id,&models.ConfirmImage{ImagePath:h.Filename,AcceptRelation:acceptRelation[0]})
		if err !=nil{//天晓得什么错误
			t.Data["json"] = HttpResponseCode{Success:false,Message:err.Error()}
		} else{
			t.Data["json"] = HttpResponseCode{Success:true,Message:"success"}
		}
		acceptRelation[0].AcTaskState = models.Task_ac_check
		_,_ = models.UpdateAcceptRelation(acceptRelation[0])
	}
	t.ServeJSON()
}


