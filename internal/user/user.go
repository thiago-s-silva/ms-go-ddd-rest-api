package user

import (
	"github.com/gin-gonic/gin"
	user_v1 "github.com/thiago-s-silva/ms-go-ddd-rest-api/internal/user/v1"
	user_v2 "github.com/thiago-s-silva/ms-go-ddd-rest-api/internal/user/v2"
)

func RegisterV1Routes(v1 *gin.RouterGroup) {
	user_v1.InitController()

	v1.GET("/user", user_v1.GetOneHandler)
}

func RegisterV2Routes(v2 *gin.RouterGroup) {
	user_v2.InitController()

	v2.GET("/user", user_v2.GetOneHandler)
}
