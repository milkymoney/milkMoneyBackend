package controllers

import (
	"apiproject/models"
	"encoding/json"
	"strconv"
	"github.com/astaxie/beego"
	"fmt"
)


type TaskController struct {
	beego.Controller
}
// @Title CreateTask
// @Description create tasks
// @Param	body		body 	models.Task	true		"body for task content"
// @Success 200 {int} models.Task.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TaskController) Post() {
	var task models.Task
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	fmt.Println(t.Ctx.Input.RequestBody)
	fmt.Println(task)
	tId,err := models.AddTask(&task)
	if err == nil{
		taskId := strconv.Itoa(tId)
		t.Data["json"] = map[string]string{"taskId": taskId}
	} else{
		t.Data["json"] = err.Error()
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