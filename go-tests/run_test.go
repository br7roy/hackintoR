package go_test

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"hackintoR/main/opts"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

func TestHelloServer(t *testing.T) {
	//url := "http://127.0.0.1:20000/balance"
	url := "http://127.0.0.1:20000/test"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	// json
	contentType := "application/json"
	data := `{"appId":"app123123123_123123","qName":"default","UserName":"realestate"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Printf("resposne:\n")
	fmt.Println(string(b))

}

func TestShExec(t *testing.T) {

	file, err := os.Open("/home/realestate_test/fth/info.sh")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	c := string(content)

	fmt.Println(c)

	cmd := exec.Command("sh", "-c", c)
	out, err := cmd.Output()
	fmt.Println(string(out), err)
	println("?????")
	println(string(out))
}

var (
	once sync.Once
)

/*func TestReadYaml(t *testing.T) {

  // https://yaml.to-go.online/
  // https://www.coder.work/article/212175
  once.Do(func() {

    filePath, err := filepath.Abs("../conf.yml")
    if err != nil {
      panic(err)
    }
    yamlFile, err := ioutil.ReadFile(filePath)

    if err != nil {
      fmt.Printf("yamlFile.Get err #%v ", err)
    }

    conf := new(opts.YmlConfig)

    err = yaml.Unmarshal(yamlFile, conf)

    if err != nil {
      panic(err)
    }

    for _, ele := range conf.Queue {
      name := ele.QueueName
      if name != "" {
        fmt.Printf("q name:%s\n", name)
      }

    }

  })

}*/

func TestReadToml(t *testing.T) {
	cfg := new(opts.TomlConfig)
	once.Do(func() {
		filePath, err := filepath.Abs("../conf.toml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
		fmt.Println(cfg)
	})

}
func TestJsonPost(t *testing.T) {
	url := "http://scheduler.paas.internal.mob.com/api/markTaskStatus"
	// 表单数据
	// json
	contentType := "application/json"
	data := `{"task_id":11697707,"status":4,"job_id":163270 ,"user_name":"futx"}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Printf("resposne:\n")
	fmt.Println(string(b))
}

func TestQ(t *testing.T) {
	opts.LoadConfig("../conf.toml")
	opts.InitDB()
	var u opts.User
	entry, e := u.QueryByEntry("admin", "222")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(entry)
	}

}
func TestSave(t *testing.T) {
	opts.LoadConfig("../conf.toml")
	opts.InitDB()
	var user opts.User
	user.QueryByEntry("dev", "333")
	uuid := opts.GenUUID()
	user.Token = uuid
	user.UpdateTokenByUser()
	//opts.DB.Model(&user).Where("id = ?", user.ID).Update("token", user.Token)
}
func TestClear(t *testing.T) {
	opts.LoadConfig("../conf.toml")
	opts.InitDB()
	var user opts.User
	user.QueryByEntry("test", "123123")

	user.ClearTokenByUser()

}

func TestFindTaskByJID(t *testing.T) {
	ids := opts.GetTaskListsByJobIds(177589, "futx")
	fmt.Printf("%v\n", ids)
}
