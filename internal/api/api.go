package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/config"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/user"
)

func Load(s *gin.Engine, cfg *config.Config) error {
	v1 := s.Group(cfg.API_V1_BASE_PATH)

	// register User routes
	user.RegisterV1Routes(v1)

	return nil
}
