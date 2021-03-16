package opts

import (
	"fmt"
	"github.com/BurntSushi/toml"
	//"gopkg.in/yaml.v2"
	//"io/ioutil"
	"path/filepath"
	"sync"
)

type TomlConfig struct {
	App struct {
		RemoteAddr string `toml:"remoteAddr"`
		RemotePort string `toml:"remotePort"`
		UsingNC    bool   `toml:"usingNC"`
		UsingSH    bool   `toml:"usingSH"`
		Test       bool   `toml:"test"`
		Schedule   struct {
			UserName string `toml:"userName"`
		} `toml:"schedule"`
		ScpInfo struct {
			Root         string `toml:"root"`
			DataSync     string `toml:"dataSync"`
			GpsConfigure string `toml:"gpsConfigure"`
			BatchDaily   string `toml:"batchDaily"`
			BatchResult  string `toml:"batchResult"`
			BatchPy      string `toml:"batchPy"`
			Month        string `toml:"month"`
		} `toml:"scpInfo"`
		Bus struct {
			Portrait []struct {
				Name string `toml:"name"`
				Type int    `toml:"type"`
			} `toml:"portrait"`
		} `toml:"bus"`
	} `toml:"app"`
	Queue []struct {
		QueueName string `toml:"queueName"`
	} `toml:"queue"`
}

var (
	Cfg  *TomlConfig
	once sync.Once
)

func LoadConfig(cPath string) *TomlConfig {
	once.Do(func() {
		Cfg = load(cPath)
	})
	return Cfg

}

func load(cPath string) *TomlConfig {
	filePath, err := filepath.Abs(cPath)
	if err != nil {
		panic(err)
	}

	conf := new(TomlConfig)

	if _, err := toml.DecodeFile(filePath, &conf); err != nil {
		panic(err)
	}

	for _, ele := range conf.Queue {
		name := ele.QueueName
		if name != "" {
			fmt.Printf("q name:%s\n", name)
		}
	}
	return conf
}

// reload config
func TriggerReload(confPath string) bool {
	fmt.Println("reload config start")
	Cfg = load(confPath)
	fmt.Println("reload config done")
	return true
}
