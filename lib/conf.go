package lib

import (
	"bytes"
	dlog "commonDemo/log"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type BaseConf struct {
	DebugMode    string    `yaml:"debug_mode"`
	TimeLocation string    `yaml:"time_location"`
	Log          LogConfig `yaml:"log"`
	Base         struct {
		DebugMode    string `yaml:"debug_mode"`
		TimeLocation string `yaml:"time_location"`
	} `yaml:"base"`
}

type LogConfig struct {
	Level string               `yaml:"log_level"`
	FW    LogConfFileWriter    `yaml:"file_writer"`
	CW    LogConfConsoleWriter `yaml:"console_writer"`
}

type LogConfFileWriter struct {
	On              bool   `yaml:"on"`
	LogPath         string `yaml:"log_path"`
	RotateLogPath   string `yaml:"rotate_log_path"`
	WFLogPath       string `yaml:"wf_log_path"`
	RotateWFLogPath string `yaml:"rotate_wf_log_path"`
}

type LogConfConsoleWriter struct {
	On    bool `yaml:"on"`
	Color bool `yaml:"color"`
}

// 全局变量
var (
	ViperConfMap map[string]*viper.Viper
	ConfBase     *BaseConf
)

func GetBaseConf() *BaseConf {
	return ConfBase
}

// 初始化配置文件
func InitViperConf() error {
	f, err := os.Open(ConfEnvPath + "/")
	if err != nil {
		return err
	}
	fileList, err := f.Readdir(1024)
	if err != nil {
		return err
	}
	for _, f0 := range fileList {
		if !f0.IsDir() {
			bts, err := ioutil.ReadFile(ConfEnvPath + "/" + f0.Name())
			if err != nil {
				return err
			}
			v := viper.New()
			v.SetConfigType("yaml")
			v.ReadConfig(bytes.NewBuffer(bts))
			pathArr := strings.Split(f0.Name(), ".") // 取出文件名
			if ViperConfMap == nil {
				ViperConfMap = make(map[string]*viper.Viper)
			}
			ViperConfMap[pathArr[0]] = v
		}
	}
	return nil
}

func InitBaseConf(path string) error {
	ConfBase := &BaseConf{}
	err := ParseConfig(path, ConfBase)
	if err != nil {
		return err
	}

	if ConfBase.DebugMode == "" {
		if ConfBase.Base.DebugMode != "" {
			ConfBase.DebugMode = ConfBase.Base.DebugMode
		} else {
			ConfBase.DebugMode = "debug"
		}
	}

	if ConfBase.TimeLocation == "" {
		if ConfBase.Base.TimeLocation != "" {
			ConfBase.TimeLocation = ConfBase.Base.TimeLocation
		} else {
			ConfBase.TimeLocation = "Asia/Chongqing"
		}
	}

	if ConfBase.Log.Level == "" {
		ConfBase.Log.Level = "trace"
	}

	// 配置日志
	logconf := dlog.LogConfig{
		Level: ConfBase.Log.Level,
		FW: dlog.LogConfFileWriter{
			On:              ConfBase.Log.FW.On,
			LogPath:         ConfBase.Log.FW.LogPath,
			RotateLogPath:   ConfBase.Log.FW.RotateLogPath,
			WFLogPath:       ConfBase.Log.FW.WFLogPath,
			RotateWFLogPath: ConfBase.Log.FW.RotateLogPath,
		},
		CW: dlog.LogConfConsoleWriter{
			On:    ConfBase.Log.CW.On,
			Color: ConfBase.Log.CW.Color,
		},
	}

	return nil
}
