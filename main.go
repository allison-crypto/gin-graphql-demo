package main

import (
	"os"
	"os/signal"
	"syscall"

	"gin-graphql-demo/config"
	"gin-graphql-demo/server"

	"github.com/sirupsen/logrus"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP, syscall.SIGQUIT)

	config.InitConfig()

	server := server.Server{}
	server.Run()

	logrus.Infoln("application started port -> ", config.System.Port)

	<-stop
	server.Shutdown()
}
