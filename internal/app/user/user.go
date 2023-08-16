package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
)

type UserDependencies struct {
	Config       *config.Config
	V1RouteGroup *gin.RouterGroup
}

func Load(d *UserDependencies) {
	d.V1RouteGroup.GET("/user/:id", GetOneHandler)
	d.V1RouteGroup.POST("/user", CreateHandler)
}
