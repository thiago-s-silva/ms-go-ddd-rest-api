package user_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-ddd-rest-api/pkg"
)

var (
	logger *pkg.Logger
)

func InitController() {
	logger = pkg.NewLogger("USER:V1")
}

func GetOneHandler(c *gin.Context) {
	GetOne(c)
}
