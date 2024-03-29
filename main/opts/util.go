package opts

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"
)

type TaskStat int

const (
	Ok   TaskStat = 4
	Fail TaskStat = 5
)

func getParams(r *http.Request) (param Params, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return param, err
	}
	if err = json.Unmarshal(body, &param); err != nil {
		return param, err
	}
	//todo basic validator
	return param, nil
}
func jsonEncode(code int, data interface{}) []byte {
	tr := new(JsonResult)
	tr.Code, tr.Data = code, data
	result, _ := json.Marshal(tr)
	return result
}

func requestDeser(r *http.Request) (param interface{}, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return param, err
	}
	if err = json.Unmarshal(body, &param); err != nil {
		return param, err
	}
	//todo basic validator
	return param, nil

}

func respSer(code int, data interface{}) []byte {
	resp, _ := json.Marshal(JsonResult{
		Code: code,
		Data: data,
	})

	return resp

}

func perform(bash string) (bool, string) {
	c := "| nc " + Cfg.App.RemoteAddr + " " + Cfg.App.RemotePort
	bash = "echo " + bash + c
	shell := ""
	if Cfg.App.UsingSH {
		shell = "sh"
	} else {
		shell = "bash"
	}
	if bash != "" {
		cmd := exec.Command(shell, "-c", bash)
		out, err := cmd.Output()
		suc := true
		if err != nil {
			suc = false
		}
		return suc, string(out)
	} else {
		return false, ""
	}

}

func localPerform(shell string) (bool, string) {
	shellType := ""
	if Cfg.App.UsingSH {
		shellType = "sh"
	} else {
		shellType = "bash"
	}
	if shell != "" {
		cmd := exec.Command(shellType, "-c", shell)
		out, err := cmd.Output()
		suc := true
		if err != nil {
			suc = false
		}
		return suc, string(out)
	} else {
		return false, ""
	}
}

func LeverTask(stat int, taskId int, jobId int, username string) (bool, string) {

	t := &Task{}
	t.UserName = username
	t.JobId = jobId
	t.TaskId = taskId
	t.Status = stat
	jsonB, err := json.Marshal(t)
	if err != nil {
		return false, ""
	}
	// todo toBe const
	return JsonPost("http://scheduler.paas.internal.mob.com/api/markTaskStatus",
		string(jsonB),
	)

}

func GetTaskListsByJobIds(jobId int, userName string) []int {
	req := SchGettasklistsbyjobidsReq{}
	req.UserName = userName
	req.JobID = jobId
	req.PageIndex = 1
	req.PageSize = 10
	jsonB, err := json.Marshal(req)
	suc, content := JsonPost("http://scheduler.paas.internal.mob.com/api/getTaskList",
		string(jsonB),
	)

	if !suc {
		return nil
	}

	resp := SchGettasklistsbyjobidsResp{}

	if err = json.Unmarshal([]byte(content), &resp); err != nil {
		return nil
	}
	var taskIds []int
	for _, ele := range resp.Data {
		taskIds = append(taskIds, ele.TaskID)
	}
	return taskIds
}

func JsonPost(url string, reqData string) (bool, string) {
	contentType := "application/json"
	fmt.Printf("request:\n")
	fmt.Println(reqData)
	resp, err := http.Post(url, contentType, strings.NewReader(reqData))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return false, "req fail"
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return false, "resp fail"
	}
	fmt.Printf("resposne:\n")
	fmt.Println(string(b))
	return true, string(b)
}

func GenUUID(seed ...uint8) string {
	v4 := uuid.NewV4()
	return v4.String()
}
