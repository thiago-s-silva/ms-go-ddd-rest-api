package user

import (
	"github.com/gin-gonic/gin"
	user_v1 "github.com/thiago-s-silva/ms-go-task-management-api/internal/user/v1"
)

func RegisterV1Routes(v1 *gin.RouterGroup) {
	user_v1.InitController()

	v1.GET("/user", user_v1.GetOneHandler)
}
