package log

import (
	"fmt"
	"sync"
)

var (
	LEVEL_FLAGS = [...]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	TARCE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

const tunnel_size_default = 1024

type Record struct {
	time  string
	code  string
	info  string
	level int
}

func (r *Record) String() string {
	return fmt.Sprintf("[%s][%s][%s] %s\n", LEVEL_FLAGS[r.level], r.time, r.code, r.info)
}

type Writer interface {
	Init() error
	Write(*Record) error
}

type Logger struct {
	writers     []Writer
	tunnel      chan *Record
	level       int
	lastTime    int64
	lastTimeStr string
	c           chan bool
	layout      string
	recordPool  *sync.Pool
}

func NewLogger() *Logger {
	if logger_default != nil && takeup == false {
		takeup = true // 默认启动标志
		return logger_default
	}

	// 初始化Logger结构体并将其返回
	l := new(Logger)
	l.writers = []Writer{}
	l.tunnel = make(chan *Record, tunnel_size_default)
	l.c = make(chan bool, 2)
	l.level = DEBUG
	l.layout = "2006/01/02 15:04:05"
	l.recordPool = &sync.Pool{New: func() interface{} {
		return &Record{}
	}}
	return l
}

// default log 全局变量
var (
	takeup         = false
	logger_default *Logger
)

func defaulyLoggerInit() {
	if takeup == false {
		logger_default = NewLogger()
	}
}
