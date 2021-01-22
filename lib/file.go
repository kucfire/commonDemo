package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	ConfEnvPath string // 配置文件夹
	ConfEnv     string // 配置环境名 比如：dev prod test
)

// 解析配置文件目录
//
// 配置文件必须放到一个文件夹中
// 如：config=conf/dev/base.yaml  ConfEnvPath=conf/dev ConfEnv=dev
// 如：config=conf/base.yaml      ConfEnvPath=conf     COnfEnv=conf
func ParseConfPath(conf string) error {
	path := strings.Split(conf, "/")                // 根据/进行分割字符串，获得单独目录的名字
	prefix := strings.Join(path[:len(path)-1], "/") // 如果分割后超过
	ConfEnvPath = prefix
	ConfEnv = path[len(path)-2]
	return nil
}

func ParseConfig(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Open Config %v failed, err: %v", path, err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Read Config %v failed, err: %v", path, err)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.ReadConfig(bytes.NewBuffer(data))
	if err := v.Unmarshal(conf); err != nil {
		return fmt.Errorf("Parse config fail, config: %v, err: %v", string(data), err)
	}
	return nil
}

// Get
func GetConfPath(filename string) string {
	return ConfEnvPath + "/" + filename + ".yaml"
}

func GetConfFilePath(filename string) string {
	return ConfEnvPath + "/" + filename
}
