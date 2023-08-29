package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/app/user"
)

type Dependencies struct {
	Server       *gin.Engine
	ServerConfig *config.ServerConfig
}

func Load(d *Dependencies) {
	v1 := d.Server.Group(d.ServerConfig.API_V1_BASE_PATH)

	// Load User Package
	user.Load(&user.Dependencies{
		Config:       d.ServerConfig,
		V1RouteGroup: v1,
	})
}
