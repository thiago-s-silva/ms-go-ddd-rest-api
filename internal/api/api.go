package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-ddd-rest-api/config"
	"github.com/thiago-s-silva/ms-go-ddd-rest-api/internal/product"
	"github.com/thiago-s-silva/ms-go-ddd-rest-api/internal/user"
)

func Load(s *gin.Engine, cfg *config.Config) error {
	v1 := s.Group(cfg.API_V1_BASE_PATH)
	v2 := s.Group(cfg.API_V2_BASE_PATH)

	// register User routes
	user.RegisterV1Routes(v1)
	user.RegisterV2Routes(v2)

	// register Product routes
	product.RegisterV1Routes(v1)

	return nil
}
