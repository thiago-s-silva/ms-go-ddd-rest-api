package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-ddd-rest-api/config"
)

func main() {
	// init Gin Gonic
	server := gin.Default()

	// init Config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// initialize Server
	server.Run(cfg.SERVER_HOST)
}
