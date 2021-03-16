package opts

type Params struct {
	AppId     string   `json:"appId"`
	QName     string   `json:"queueName"`
	UserName  string   `json:"userName"`
	Date1     string   `json:"date1"`
	Date2     string   `json:"date2"`
	Fc        bool     `json:"fc"`
	City      string   `json:"city"`
	StartTime string   `json:"startTime"`
	EndTime   string   `json:"endTime"`
	IsPre     bool     `json:"isPre"`
	PrdEnv    bool     `json:"prdEnv"`
	Types     []uint16 `json:"types"`
	Uuids     []struct {
		Value string `json:"value"`
	} `json:"uuids"`
	Cmd     string `json:"cmd"`
	Secreat string `json:"secreat"`
	ClsName string `json:"clsName"`
	JobId   int    `json:"jobId"`
	TaskId  []int  `json:"taskId"`
	Stat    int    `json:"stat"`
}

type JsonResult struct {
	Code int         `json:"code"` // 0：success 1:faild
	Data interface{} `json:"data"` // 返回的data
}

type Task struct {
	TaskId   int    `json:"task_id"`
	Status   int    `json:"status"`
	JobId    int    `json:"job_id"`
	UserName string `json:"user_name"`
}
