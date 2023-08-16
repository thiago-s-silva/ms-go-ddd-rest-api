package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

type IUserService interface {
	Create(payload *UserRequestPayload)
	GetOne(id string)
}

type UserService struct {
	context *gin.Context
	log     *pkg.Logger
}

func (u *UserService) Create(payload *UserRequestPayload) {
	u.log.Info("created new user")
	pkg.OnSuccessWithData(u.context, payload)
}

func (u *UserService) GetOne(id string) {
	u.log.Info("get one user")
	pkg.OnSuccessWithData(u.context, id)
}
