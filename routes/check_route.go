package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func mountCheckRoutes(checkRoute *gin.RouterGroup) {
	checkRoute.GET("/", func(ctx *gin.Context) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "localhost"
		}

		ctx.JSON(http.StatusOK, gin.H{
			"msg":       "OK",
			"timestamp": time.Now().Format("2006-01-02 15:04:05.000"),
			"hostname":  hostname,
		})
	})
}
