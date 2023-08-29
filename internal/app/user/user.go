package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
)

type Dependencies struct {
	Config       *config.ServerConfig
	V1RouteGroup *gin.RouterGroup
}

func Load(d *Dependencies) {
	d.V1RouteGroup.GET("/user/:id", GetOneHandler)
	d.V1RouteGroup.POST("/user", CreateHandler)
}
