package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/app/user"
)

type Dependencies struct {
	Server *gin.Engine
	Config *config.Config
}

func (d *Dependencies) Load() {
	v1 := d.Server.Group(d.Config.API_V1_BASE_PATH)

	// Load User Package
	user.Load(&user.UserDependencies{
		Config:       d.Config,
		V1RouteGroup: v1,
	})
}
