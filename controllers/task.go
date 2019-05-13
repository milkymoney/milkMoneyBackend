package controllers

import (
	"github.com/milkymoney/milkMoneyBackend/models"
	"encoding/json"
	"strconv"
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

// @Title CreateTask
// @Description specific user create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]

func (t *TaskController) Post() {
	fmt.Println("Create Task")
	var task models.Task
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
	} else{
		task.Userid = user.Id
		fmt.Println(task)
		tId,err := models.AddTask(&task)
		if err == nil{
			taskId := strconv.Itoa(tId)
			t.Data["json"] = map[string]string{"taskId": taskId}
		} else{
			t.Data["json"] = err.Error()
		}
	}
	t.ServeJSON()
}


// @Title GetTask
// @Description get task by taskId
// @Param	taskid		path 	string	true		"the key"
// @Success 200 {object} models.Task
// @Failure 403 :taskid is empty
// @router /:taskid [get]
func (t *TaskController) Get() {
	tid := t.GetString(":taskid")
	if tid != "" {
		taskId,err := strconv.Atoi(tid)
		if err == nil{
			task, err := models.GetTask(taskId)
			if err != nil {
				t.Data["json"] = err.Error()
			} else {
				t.Data["json"] = task
			}
		} else{
			t.Data["json"] = err.Error()
		}

	}
	t.ServeJSON()
}
// @Title GetAllTask
// @Description get task by taskId
// @Param	openid		header 	string	true		"user's id from wx"
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

// @Title UpdateTask
// @Description update the task
// @Param	taskid		path 	string	true		"The taskId you want to update"
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {object} models.Task
// @Failure 403 :taskid is not int
// @router /:taskid [put]
func (t *TaskController) Put() {
	tid := t.GetString(":taskid")
	if tid != "" {
		var task models.Task
		json.Unmarshal(t.Ctx.Input.RequestBody, &task)
		taskId,err := strconv.Atoi(tid)
		if err == nil{
			tt, err := models.UpdateTask(taskId, &task)
			if err != nil {
				t.Data["json"] = err.Error()
			} else {
				t.Data["json"] = tt
			}
		} else{
			t.Data["json"] = err.Error()
		}

	}
	t.ServeJSON()
}

// @Title DeleteTask
// @Description delete the task
// @Param	taskid		path 	string	true		"The taskid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 taskid is empty
// @router /:taskid [delete]
func (t *TaskController) Delete() {
	tid := t.GetString(":taskid")
	taskId,err := strconv.Atoi(tid)
	if err ==nil{
		models.DeleteTask(taskId)
		t.Data["json"] = "delete success!"
	} else{
		t.Data["json"] = err.Error()
	}

	t.ServeJSON()
}
//注：暂时不检查是否达到最大值，只是进行处理
// @Title AcceptTask
// @Description user accept the task
// @Param	taskid		path 	string	true		"The taskid you want to accept"
// @Success 200 {string} accept success!
// @Failure 403 taskid is empty
// @router /accept/:taskid [put]
func (t *TaskController) AcceptTask(){
	user,err := Auth(&t.Controller)
	if err != nil{
		t.Data["json"] = err.Error()
	} else{
		tid := t.GetString(":taskid")
		taskId,err := strconv.Atoi(tid)
		if err != nil{
			t.Data["json"] = err.Error()
		}else{
			_,err = models.CreateNewAcRelById(user.Id,taskId,"20190513")
			if err == nil{
				t.Data["json"] = "accept success!"
			} else{
				t.Data["json"] = err.Error()
			}
		}
	}
	t.ServeJSON()
}
