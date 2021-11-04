package config

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var System *SystemConfig

func InitConfig() {
	port, err := strconv.Atoi(os.Getenv("PROT"))
	if err != nil {
		port = 8000
	}
	stage := os.Getenv("STAGE")
	mode := gin.ReleaseMode
	if stage == "local" {
		mode = gin.DebugMode
	}
	System = &SystemConfig{
		Stage: stage,
		Port:  port,
		Mode:  mode,
	}
}

type SystemConfig struct {
	Stage string
	Port  int
	Mode  string
}
