package app

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/app/user"
	"gorm.io/gorm"
)

type Dependencies struct {
	Server *gin.Engine
	Config *config.Config
	Db     *gorm.DB
}

func Load(d *Dependencies) error {
	v1 := d.Server.Group(d.Config.ServerConfig.API_V1_BASE_PATH)

	// Load User Package
	err := user.Load(&user.Dependencies{
		Config:       d.Config,
		V1RouteGroup: v1,
		Db:           d.Db,
	})
	if err != nil {
		return err
	}

	return nil
}
