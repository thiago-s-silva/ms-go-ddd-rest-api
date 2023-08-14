package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/api"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

func main() {
	// init Logger
	logger := pkg.NewLogger("main.go")

	// init Gin Gonic
	server := gin.Default()

	// init Config
	cfg, err := config.Load()
	if err != nil {
		logger.Errorf("error when tried to load config package: %v", err)
		panic(err)
	}

	// init API
	api.Load(server, cfg)

	// initialize Server
	server.Run(cfg.SERVER_HOST)
}
