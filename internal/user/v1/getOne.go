package user_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-ddd-rest-api/pkg"
)

func GetOne(c *gin.Context) {
	logger.Info("got request from v1")
	pkg.OnSuccess(c, "working from v1")
}
