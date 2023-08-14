package product

import (
	"github.com/gin-gonic/gin"
	product_v1 "github.com/thiago-s-silva/ms-go-ddd-rest-api/internal/product/v1"
)

func RegisterV1Routes(v1 *gin.RouterGroup) {
	product_v1.InitController()

	v1.GET("/product", product_v1.GetOneHandler)
}
