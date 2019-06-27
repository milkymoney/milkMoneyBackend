package test

import (
	"net/http"
	"net/http/httptest"
	// "net/url"
	"testing"
	"runtime"
	"path/filepath"
	"encoding/json"
	"io/ioutil"
	_ "github.com/milkymoney/milkMoneyBackend/routers"
	// "github.com/milkymoney/milkMoneyBackend/controllers"
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
// func TestGetAllTask(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/task/?userId=2", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var tasks []models.Task
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &tasks)
// 	logs.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
// 	// logs.Trace(tasks)
	
// 	// logs.Trace(tasks[3])
// 	// logs.Trace(tasks[0].Userid)
	
// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 	        Convey("Status Code Should Be 200", func() {
// 	                So(w.Code, ShouldEqual, 200)
// 			})
// 			Convey("Task Id should Be 62", func() {
// 				So(tasks[1].Id, ShouldEqual, 62)
// 			})
// 			Convey("User Id should Be 13", func() {
// 				So(tasks[1].Userid, ShouldEqual, 13)
// 			})
// 			Convey("Reward should Be  1", func() {
// 				So(tasks[1].Reward, ShouldEqual, 1)
// 			})
// 			Convey("Max Accept should Be 1", func() {
// 				So(tasks[1].MaxAccept, ShouldEqual, 1)
// 			})
// 	})
// }

// func TestOwnTask(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/task/publisher?userId=11", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var tasks []models.Task
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &tasks)
// 	logs.Trace(tasks)
// 	// logs.Trace(tasks[1].Type)
// 	// logs.Trace(tasks[1].Description)

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 	        Convey("Status Code Should Be 200", func() {
// 	                So(w.Code, ShouldEqual, 200)
// 			})
// 			Convey("Task Id should Be 56", func() {
// 				So(tasks[0].Id, ShouldEqual, 56)
// 			})
// 			Convey("User Id should Be 11", func() {
// 				So(tasks[0].Userid, ShouldEqual, 11)
// 			})
// 			Convey("Reward should Be  1", func() {
// 				So(tasks[0].Reward, ShouldEqual, 1)
// 			})
// 			Convey("Max Accept should Be 1", func() {
// 				So(tasks[0].MaxAccept, ShouldEqual, 1)
// 			})
// 	})
// }

// func TestOwnTaskId(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/task/publisher/34?userId=2", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var tasks models.Task
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &tasks)
// 	logs.Trace(tasks.Id)

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 			Convey("Status Code Should Be 200", func() {
// 				So(w.Code, ShouldEqual, 200)
// 			})
// 			Convey("Task Id should Be 34", func() {
// 				So(tasks.Id, ShouldEqual, 34)
// 			})
// 			Convey("User Id should Be 2", func() {
// 				So(tasks.Userid, ShouldEqual, 2)
// 			})
// 			Convey("Reward should Be  53", func() {
// 				So(tasks.Reward, ShouldEqual, 3)
// 			})
// 			Convey("Max Accept should Be 10", func() {
// 				So(tasks.MaxAccept, ShouldEqual, 10)
// 			})
// 	})
// }

func TestAllInfomation(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/publisher/confirm/43?userId=2", nil)
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
		// Convey("Task Id should Be 35", func() {
		// 	So(tasks[0].Id, ShouldEqual, 35)
		// })
		// Convey("User Id should Be 3", func() {
		// 	So(tasks[0].Userid, ShouldEqual, 3)
		// })
		// Convey("Reward should Be  5", func() {
		// 	So(tasks[0].Reward, ShouldEqual, 5)
		// })
		// Convey("Max Accept should Be 5", func() {
		// 	So(tasks[0].MaxAccept, ShouldEqual, 5)
		// })
	})
}

// func TestTaskRecipient(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/task/recipient/?userId=3", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var tasks []models.Task
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &tasks)
// 	logs.Trace(tasks[0].Id)
// 	logs.Trace(tasks[0].Userid)

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 		Convey("Status Code Should Be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("Task Id Should Be 2", func() {
// 			So(tasks[0].Id,ShouldEqual, 2)
// 		})
// 		Convey("UserID Who Release The Task Should Be 2", func() {
// 			So(tasks[0].Userid,ShouldEqual, 2)
// 		})
// 	})
// }

// func TestTaskRecipientById(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/task/recipient/2?userId=3", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var tasks models.Task
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &tasks)
// 	logs.Trace(tasks.Id)
// 	logs.Trace(tasks.MaxAccept)

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 		Convey("Status Code Should Be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("Task Id Should Be 2", func() {
// 			So(tasks.Id,ShouldEqual, 2)
// 		})
// 		Convey("MaxAccept Should Be 5", func() {
// 			So(tasks.MaxAccept,ShouldEqual, 5)
// 		})
// 	})
// }


// func TestTaskRecipientFinished(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/task/recipient/settleup/2?userId=3", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var tasks controllers.AcceptorCheckFinishCodeResponse
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &tasks)
// 	logs.Trace(tasks.Proves)

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 		Convey("Status Code Should Be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("Prove URL Should Be fadfas", func() {
// 			So(tasks.Proves[0], ShouldEqual, "fadfas")
// 		})
// 	})
// }

// func TestUser(t *testing.T) {
// 	r, _ := http.NewRequest("GET", "/v1/user?userId=2", nil)
// 	w := httptest.NewRecorder()
// 	beego.BeeApp.Handlers.ServeHTTP(w, r)

// 	var users models.User
// 	readBuf,_ := ioutil.ReadAll(w.Body)
// 	json.Unmarshal(readBuf, &users)
// 	logs.Trace(users.Id)
// 	logs.Trace(users.Balance)

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 		Convey("Status Code Should Be 200", func() {
// 			So(w.Code, ShouldEqual, 200)
// 		})
// 		Convey("Task Id Should Be 2", func() {
// 			So(users.Id,ShouldEqual, 2)
// 		})
// 		Convey("MaxAccept Should Be 24", func() {
// 			So(users.Balance,ShouldEqual, 24)
// 		})
// 	})
// }

// func TestPostTask (t *testing.T) {
// 	// postData := []byte(`{"reward": 5, "maxaccept": 5}`)
// 	// r, _ := http.NewRequest("POST", "/v1/task/publisher?userId=3", nil)
	
// 	response, _ := http.PostForm("https://www.wtysysu.cn:10443/v1/task/publisher?userId=3", url.Values{"reward":{"10"}, "maxaccept":{"5"}})
// 	body, _ := ioutil.ReadAll(response.Body)
	
// 	var tasks models.Task 
// 	json.Unmarshal(body,&tasks)
// 	logs.Trace(tasks)
// 	logs.Trace(string(body))

// 	Convey("Subject: Test Station Endpoint\n", t, func() {
// 		Convey("Task Id Should Be 2", func() {
// 			So(tasks.Id,ShouldEqual, 2)
// 		})
// 		Convey("MaxAccept Should Be 5", func() {
// 			So(tasks.MaxAccept,ShouldEqual, 5)
// 		})
// 	})
// }
