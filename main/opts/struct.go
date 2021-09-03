package opts

const uploadPath = "/opt/flink/conf"

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
	Cmd         string     `json:"cmd"`
	Secreat     string     `json:"secreat"`
	ClsName     string     `json:"clsName"`
	JobId       int        `json:"jobId"`
	TaskId      []int      `json:"taskId"`
	JobIds      []int      `json:"jobIds"`
	SchUserName string     `json:"schUserName"`
	Stat        int        `json:"stat"`
	LoginName   string     `json:"loginName"`
	Password    string     `json:"password"`
	Token       string     `json:"token"`
	ConfParam   TomlConfig `json:"confParam"`
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

// 调度系统请求参数
type SchGettasklistsbyjobidsReq struct {
	StartBeginTime interface{} `json:"start_begin_time"`
	Content        interface{} `json:"content"`
	JobID          int         `json:"job_id"`
	PageIndex      int         `json:"page_index"`
	PageSize       int         `json:"page_size"`
	StartDataTime  interface{} `json:"start_data_time"`
	TaskStatus     interface{} `json:"task_status"`
	AppID          interface{} `json:"app_id"`
	JobType        interface{} `json:"job_type"`
	OrderColumn    struct {
		Key  interface{} `json:"key"`
		Type interface{} `json:"type"`
	} `json:"order_column"`
	UserName     string      `json:"user_name"`
	EndBeginTime interface{} `json:"end_begin_time"`
	EndDataTime  interface{} `json:"end_data_time"`
}

// 调度系统返回参数
type SchGettasklistsbyjobidsResp struct {
	Msg   string `json:"msg"`
	Total int    `json:"total"`
	Data  []struct {
		JobType           int    `json:"job_type"`
		EndTime           string `json:"end_time"`
		BeginTime         string `json:"begin_time"`
		ExecuteServerHost string `json:"execute_server_host"`
		TaskID            int    `json:"task_id"`
		Priority          int    `json:"priority"`
		ExecuteServerName string `json:"execute_server_name"`
		Duration          int    `json:"duration"`
		IsDisableRun      int    `json:"is_disable_run"`
		ExecuteTimes      int    `json:"execute_times"`
		JobName           string `json:"job_name"`
		DataTime          string `json:"data_time"`
		FirstTime         string `json:"first_time"`
		JobID             int    `json:"job_id"`
		NorminalTime      string `json:"norminal_time"`
		ExecuteServerID   int    `json:"execute_server_id"`
		Status            int    `json:"status"`
	} `json:"data"`
	Status int `json:"status"`
}
