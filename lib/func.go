package lib

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// 全局函数
var (
	LocalIP    = net.ParseIP("127.0.0.1")
	TimeFormat = "2006-01-02 15:04:05"
)

func Init(path string) error {
	return InitModule(path, []string{"base", "mysql", "redis"})
}

func InitModule(path string, configName []string) error {
	conf := flag.String("config", path, "input config file path like ./conf/dev//")
	flag.Parse()
	if *conf == "" {
		flag.Usage()
		os.Exit(1)
	}

	log.Println("---------------------------------------------------------------")
	log.Printf("[INFO] config=%s\n", *conf)
	log.Println("[INFO] start loading resource")

	// 设置ip信息，优先设置便于日志打印
	ips := GetLocalIPs()
	if len(ips) > 0 {
		LocalIP = ips[0]
	}

	// 解析配置文件目录
	if err := ParseConfPath(*conf); err != nil {
		return err
	}

	//初始化配置文件
	if err := InitViperConf(); err != nil {
		return err
	}

	// 加载base配置
	if InArrayString("base", configName) {
		if err := InitBaseConf(GetConfPath("base")); err != nil {
			fmt.Printf("[ERROR] %s%s\n", time.Now().Format(TimeFormat), " InitBaseConf:"+err.Error())
		}
	}
	return nil
}

func GetLocalIPs() (ips []net.IP) {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}

	for _, address := range interfaceAddrs {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			ips = append(ips, ipNet.IP)
		}
	}
	return ips
}

func InArrayString(s string, arr []string) bool {
	for _, value := range arr {
		if s == value {
			return true
		}
	}
	return false
}
