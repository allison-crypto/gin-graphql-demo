package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"gin-graphql-demo/config"
	"gin-graphql-demo/server"

	"github.com/sirupsen/logrus"
)

func main() {
	var configPath string

	//监听程序退出信号
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)

	//解析flag
	flag.StringVar(&configPath, "c", "./", "config file path")
	flag.Parse()

	config.InitConfig(configPath)
	logrus.Infof("%+v \n", config.C)

	server := server.Server{Config: config.C.Server}
	server.Run()

	logrus.Infoln("application started port -> ", config.C.Server.Port)

	<-stop
	server.Shutdown()
}
