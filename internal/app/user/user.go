package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"gorm.io/gorm"
)

type Dependencies struct {
	Config       *config.ServerConfig
	V1RouteGroup *gin.RouterGroup
	Db           *gorm.DB
}

func Load(d *Dependencies) error {
	// auto migrate User table
	if err := d.Db.AutoMigrate(&User{}); err != nil {
		return err
	}

	// register endpoints
	d.V1RouteGroup.GET("/user/:id", GetOneHandler)
	d.V1RouteGroup.POST("/user", CreateHandler)

	return nil
}
