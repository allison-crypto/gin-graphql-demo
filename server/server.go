package server

import (
	"context"
	"fmt"
	"gin-graphql-demo/config"
	"gin-graphql-demo/routes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Server with gin framework
type Server struct {
	httpServer *http.Server
}

// Run to start a gin server
func (server *Server) Run() {
	gin.SetMode(config.System.Mode)
	app := gin.Default()
	routes.MountRoutes(app)

	server.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", config.System.Port),
		Handler: app,
	}

	go func() {
		if err := server.httpServer.ListenAndServe(); err != nil {
			logrus.Fatalln("start server error: ", err.Error())
		}
	}()
}

// Shutdown to stop a gin server
func (server *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := server.httpServer.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown:", err)
	}
	logrus.Info("application stopped")
}
