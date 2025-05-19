package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/owenliang/go-push/logic"
)

var (
	confFile string // 配置文件路径
)

func initArgs() {
	flag.StringVar(&confFile, "config", "./logic.json", "where logic.json is.")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化环境
	initArgs()
	initEnv()
	// 初始化配置文件
	if err = logic.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 初始化统计
	if err = logic.InitStats(); err != nil {
		goto ERR
	}
	// 初始化网关连接管理器
	if err = logic.InitGateConnMgr(); err != nil {
		goto ERR
	}
	// 初始化服务
	if err = logic.InitService(); err != nil {
		goto ERR
	}

	// 主循环 死循环不让退出
	for {
		time.Sleep(1 * time.Second)
	}

	os.Exit(0)

ERR:
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
