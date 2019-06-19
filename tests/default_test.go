package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	_ "github.com/milkymoney/milkMoneyBackend/routers"
	"github.com/milkymoney/milkMoneyBackend/controllers"
	"github.com/milkymoney/milkMoneyBackend/models"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/astaxie/beego/logs"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

//TestGet is a sample to run an endpoint test
func TestGetAllTask(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/?userId=2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks []models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	logs.Trace(tasks)
	
	logs.Trace(tasks[0].Id)
	logs.Trace(tasks[0].Userid)
	
	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	})
}

func TestOwnTask(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/publisher?userId=2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks []models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks)
	// logs.Trace(tasks[1].Type)
	// logs.Trace(tasks[1].Description)

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	})
}

func TestOwnTaskId(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/publisher/2?userId=2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks.Id)

	Convey("Subject: Test Station Endpoint\n", t, func() {
			Convey("Status Code Should Be 200", func() {
				So(w.Code, ShouldEqual, 200)
			})
			Convey("Task Id should Be 2", func() {
				So(tasks.Id, ShouldEqual, 2)
			})
	})
}

func TestAllInfomation(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/publisher/confirm/2?userId=3", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks []models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestTaskRecipient(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/recipient/?userId=3", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks []models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks[0].Id)
	logs.Trace(tasks[0].Userid)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Task Id Should Be 2", func() {
			So(tasks[0].Id,ShouldEqual, 2)
		})
		Convey("UserID Who Release The Task Should Be 2", func() {
			So(tasks[0].Userid,ShouldEqual, 2)
		})
	})
}

func TestTaskRecipientById(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/recipient/2?userId=3", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks.Id)
	logs.Trace(tasks.MaxAccept)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Task Id Should Be 2", func() {
			So(tasks.Id,ShouldEqual, 2)
		})
		Convey("MaxAccept Should Be 5", func() {
			So(tasks.MaxAccept,ShouldEqual, 5)
		})
	})
}


func TestTaskRecipientFinished(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/recipient/settleup/2?userId=3", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks controllers.AcceptorCheckFinishCodeResponse
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks.Proves)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Prove URL Should Be fadfas", func() {
			So(tasks.Proves[0], ShouldEqual, "fadfas")
		})
	})
}

func TestUser(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/user?userId=3", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace(tasks.Id)
	logs.Trace(tasks.MaxAccept)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("Task Id Should Be 2", func() {
			So(tasks.Id,ShouldEqual, 2)
		})
		Convey("MaxAccept Should Be 5", func() {
			So(tasks.MaxAccept,ShouldEqual, 5)
		})
	})
}
