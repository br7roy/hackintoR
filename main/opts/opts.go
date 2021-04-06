package opts

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type InitOptions struct {
	Mux *http.ServeMux
}
type Opts struct {
	mux *http.ServeMux
}

func WakeOpts(initOpts InitOptions) *Opts {
	opt := new(Opts)
	opt.mux = initOpts.Mux
	opt.mallocRegHandler()
	return opt
}
func (se *Opts) mallocRegHandler() {
	se.mux.HandleFunc("/go-test", se.test)
	se.mux.HandleFunc("/balance", se.balance)
	se.mux.HandleFunc("/qqq", se.queryQueue)
	se.mux.HandleFunc("/py", se.py)
	se.mux.HandleFunc("/result", se.result)
	se.mux.HandleFunc("/daily", se.daily)
	se.mux.HandleFunc("/mon", se.mon)
	se.mux.HandleFunc("/ds", se.ds)
	se.mux.HandleFunc("/gps", se.gps)
	se.mux.HandleFunc("/hxtp", se.hxType)
	se.mux.HandleFunc("/definePoly", se.definePoly)
	se.mux.HandleFunc("/ssupload", se.ssupload)
	se.mux.HandleFunc("/localshell", se.localShell)
	se.mux.HandleFunc("/bt", se.batchTask)
	se.mux.HandleFunc("/login", se.login)
	se.mux.HandleFunc("/usrinfo", se.usrinfo)
	se.mux.HandleFunc("/logout", se.logout)
	se.mux.HandleFunc("/bj", se.batchJob)

}

// just a test Handler
func (se *Opts) test(w http.ResponseWriter, r *http.Request) {
	params, err := requestDeser(r)
	if err != nil {
		w.Write(respSer(1, map[string]interface{}{"error": "testerror"}))
		return
	}
	res, _ := json.Marshal(params)
	fmt.Println(string(res))
	w.Write(respSer(0, params))

}

// 定点围栏
func (se *Opts) definePoly(w http.ResponseWriter, r *http.Request) {

	params, done := valid(w, r)
	if done {
		return
	}

	w.WriteHeader(200)
	w.Write(respSer(0, params))

}

// move the application to specify queue
func (se *Opts) balance(w http.ResponseWriter, r *http.Request) {

	params, done := valid(w, r)
	if done {
		return
	}
	replacer := strings.NewReplacer("root.", "")
	qname := replacer.Replace(params.QName)

	command := "HADOOP_USER_NAME=" + params.UserName + "  yarn application  -queue " + qname + "  --movetoqueue " + params.AppId

	suc, msg := perform(command)

	if suc {
		fmt.Println("exec good")
		w.Write(respSer(0, "ok"))
	} else {
		fmt.Println("exec fail")
		w.Write(respSer(1, msg))
	}
	w.WriteHeader(200)

}

// list all the queue
func (se *Opts) queryQueue(w http.ResponseWriter, r *http.Request) {
	var ques []string

	for _, ele := range Cfg.Queue {
		name := ele.QueueName
		if name != "" {
			ques = append(ques, name)
		}
	}
	w.WriteHeader(200)
	w.Write(respSer(0, ques))
}

// list all the portrait type
func (se *Opts) hxType(w http.ResponseWriter, r *http.Request) {

	var infos []interface{}

	for _, ele := range Cfg.App.Bus.Portrait {
		infos = append(infos, ele)
	}

	//var infos []map[string]string
	//
	//for _,ele := range Cfg.App.Bus.Portrait {
	//  infos = append(infos,map[string]string{ele.Name: ele.Type})
	//}

	w.WriteHeader(200)
	w.Write(respSer(0, infos))
}
func (se *Opts) py(w http.ResponseWriter, r *http.Request) {
	params, done := valid(w, r)
	if done {
		return
	}
	performByCmdName(Cfg.App.ScpInfo.BatchPy, params, w)

}

func (se *Opts) result(w http.ResponseWriter, r *http.Request) {
	params, done := valid(w, r)
	if done {
		return
	}

	performByCmdName(Cfg.App.ScpInfo.BatchResult, params, w)

}

func (se *Opts) daily(w http.ResponseWriter, r *http.Request) {
	_, done := valid(w, r)
	if done {
		return
	}
	//performByCmdName(Cfg.App.ScpInfo.BatchDaily, params, w)
	w.WriteHeader(200)
	w.Write(respSer(0, "ok"))
}

func (se *Opts) mon(w http.ResponseWriter, r *http.Request) {
	params, done := valid(w, r)
	if done {
		return
	}

	performByCmdName(Cfg.App.ScpInfo.Month, params, w)
}

func (se *Opts) ds(w http.ResponseWriter, r *http.Request) {
	params, done := valid(w, r)
	if done {
		return
	}

	performByCmdName(Cfg.App.ScpInfo.DataSync, params, w)

}

func (se *Opts) gps(w http.ResponseWriter, r *http.Request) {

	params, done := valid(w, r)
	if done {
		return
	}
	performByCmdName(Cfg.App.ScpInfo.GpsConfigure, params, w)
}

func (se *Opts) ssupload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(100)
	mForm := r.MultipartForm
	var jafileName string
	for k, _ := range mForm.File {
		// k = files's key
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			fmt.Printf("invoke form error:%s\n", err)
			return
		}
		defer file.Close()
		fmt.Printf("the uploaded file: name[%s] , size[%d], header[%#v]\n",
			fileHeader.Filename, fileHeader.Size, fileHeader.Header)
		jafileName = uploadPath + "/" + fileHeader.Filename
		out, err := os.Create(jafileName)
		if err != nil {
			fmt.Printf("create file error:%s\n", err)
			w.Write(respSer(1, "create file error"))
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Printf("copy file error:%s\n", err)
			w.Write(respSer(1, "copy file error"))
			return
		}
		fmt.Println("upload success")
	}
	value := mForm.Value
	if value == nil {
		w.Write(respSer(1, "value is nil "))
		return
	}
	appId := value["appId"]

	clsName := value["clsName"]

	Cmd := "/opt/flink/bin/flink run -c " + clsName[0] + " -yid " + appId[0] + "  " + jafileName

	//  tmpsh := `#!/bin/env bash
	//` + Cmd + `
	//echo "submit ok"
	//`
	//  fmt.Printf("on fire: %s\r\n", tmpsh)
	//
	//
	//  b := []byte(tmpsh)
	//  e1 := ioutil.WriteFile(uploadPath+"/tmpsh", b, 0755)
	//  if e1 != nil {
	//    fmt.Println("write file error")
	//    panic(e1)
	//  }

	cmd := exec.Command("bash", "-c", Cmd)
	bt, er := cmd.Output()
	if er != nil {
		fmt.Errorf("erro :%s\r\n", er)
		if bt != nil {
			fmt.Println("error msg :" + string(bt))
		}
		w.Write(respSer(1, "提交失败"))
		return
	}

	w.Write(respSer(0, "提交ok,去UI看看吧"))

}

func (se *Opts) localShell(w http.ResponseWriter, r *http.Request) {
	params, done := valid(w, r)
	if done {
		return
	}
	fmt.Println(params)
	if params.Secreat != "whoami" {
		w.Write(respSer(1, "You have no permission to execute!"))
		return
	}
	suc, msg := localPerform(params.Cmd)
	if suc {
		fmt.Println("exec good")
		w.Write(respSer(0, msg))
	} else {
		fmt.Println("exec fail")
		w.Write(respSer(1, "exec fail"))
	}
	w.WriteHeader(200)
}

func (se *Opts) batchTask(w http.ResponseWriter, r *http.Request) {

	params, done := valid(w, r)
	if done {
		return
	}
	fmt.Println(params)
	for _, elem := range params.TaskId {
		LeverTask(params.Stat, elem, params.JobId, params.SchUserName)
	}
	w.Write(respSer(0, "任务提交成功"))
}

func (se *Opts) login(w http.ResponseWriter, r *http.Request) {
	p, done := valid(w, r)
	if done {
		return
	}
	var user User
	entry, err := user.QueryByEntry(p.LoginName, p.Password)
	//token
	if err != nil {
		w.Write(respSer(1, "密码不正确"))
	} else {
		uuid := GenUUID()
		entry.Token = uuid
		entry.UpdateTokenByUser()
		w.Write(respSer(0, entry))
	}

}

func (se *Opts) usrinfo(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	token := vars["token"][0]
	var user User
	entry, err := user.QueryByToken(token)
	//token
	if err != nil {
		w.Write(respSer(1, "重新登录试试"))
	} else {
		w.Write(respSer(0, entry))
	}

}

func (se *Opts) logout(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query()["token"][0]
	var user User
	_, err := user.QueryByToken(token)
	if err == nil {
		user.ClearTokenByUser()
	}
	w.Write(respSer(0, "clear!"))

}

func (se *Opts) batchJob(w http.ResponseWriter, r *http.Request) {
	params, done := valid(w, r)
	if done {
		return
	}
	fmt.Println(params)
	for _, jobId := range params.JobIds {
		tIds := GetTaskListsByJobIds(jobId, params.SchUserName)
		for _, taskId := range tIds {
			LeverTask(params.Stat, taskId, jobId, params.SchUserName)
		}
	}
	w.Write(respSer(0, "任务提交成功"))
}

func valid(w http.ResponseWriter, r *http.Request) (Params, bool) {
	params, err := getParams(r)
	if err != nil {
		w.Write(jsonEncode(1, map[string]interface{}{"error": "param valid error"}))
		return Params{}, true
	}
	return params, false
}

func performByCmdName(cmdName string, params Params, w http.ResponseWriter) {
	dir := Cfg.App.ScpInfo.Root
	var cmd string
	switch cmdName {
	case Cfg.App.ScpInfo.BatchPy:
		cmd = dir + "/" + Cfg.App.ScpInfo.BatchPy + "  " + params.StartTime + "  " + params.EndTime + "  " + strconv.FormatBool(params.IsPre)
	case Cfg.App.ScpInfo.BatchResult:
		cmd = dir + "/" + Cfg.App.ScpInfo.BatchResult + "  " + params.StartTime + "  " + params.EndTime
	case Cfg.App.ScpInfo.BatchDaily:
		cmd = dir + "/" + Cfg.App.ScpInfo.BatchDaily + "  " + params.StartTime + "  " + params.EndTime
	case Cfg.App.ScpInfo.Month:
		cmd = dir + "/" + Cfg.App.ScpInfo.Month + "  " + params.StartTime + "  " + params.EndTime
	case Cfg.App.ScpInfo.DataSync:
		cmd = dir + "/" + Cfg.App.ScpInfo.DataSync + "  " + params.StartTime + "  " + params.EndTime
	case Cfg.App.ScpInfo.GpsConfigure:
		cmd = dir + "/" + Cfg.App.ScpInfo.GpsConfigure + "  " + params.StartTime + "  " + params.EndTime
	}
	suc, msg := perform(cmd)
	if suc {
		fmt.Println("exec good")
		w.Write(respSer(0, "Submit ok"))
	} else {
		fmt.Println("exec fail")
		w.Write(respSer(1, msg))
	}
	w.WriteHeader(200)

}
