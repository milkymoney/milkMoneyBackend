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

// TestGet is a sample to run an endpoint test
func TestGetAllTask(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/task/?userId=2", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var tasks []models.Task
	readBuf,_ := ioutil.ReadAll(w.Body)
	json.Unmarshal(readBuf, &tasks)
	logs.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
	logs.Trace(tasks)
	/*
	logs.Trace(tasks[0].Id)
	logs.Trace(tasks[0].Userid)
*/
	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	})
}

