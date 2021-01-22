package log

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

func SetupLogInstanceWithCOnf(lc LogConfig, logger *Logger) (err error) {
	if lc.FW.On {
		if len(lc.FW.LogPath) > 0 {
			w := NewFileWriter()
			w.SetFileName(lc.FW.LogPath)
		}

		if len(lc.FW.WFLogPath) > 0 {

		}
	}
}

func SetupDefaultLogWithConf(lc LogConfig) (err error) {
	defaulyLoggerInit()
	return SetupLogInstanceWithConf(lc, logger_default)
}
