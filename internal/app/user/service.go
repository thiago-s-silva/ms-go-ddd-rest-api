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
}

func (u *UserService) Create(payload *UserRequestPayload) {
	pkg.OnSuccessWithData(u.context, payload)
}

func (u *UserService) GetOne(id string) {
	pkg.OnSuccessWithData(u.context, id)
}
