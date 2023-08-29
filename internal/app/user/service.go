package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

type IUserService interface {
	Create(payload *UserRequestPayload)
	GetAll()
}

type UserService struct{}

func (u *UserService) Create(ctx *gin.Context, log *pkg.Logger, p *UserRequestPayload) {
	user := User{
		Name:     p.Name,
		Role:     p.Role,
		Email:    p.Email,
		Password: p.Password,
	}

	if err := repository.Create(&user); err != nil {
		log.Errorf("failed when tried to create new user: %s", err.Error())
		pkg.OnError(ctx, err.Error())
	}

	log.Info("created new user")
	pkg.OnSuccessWithData(ctx, user)
}

func (u *UserService) GetAll(ctx *gin.Context, log *pkg.Logger) {
	users, err := repository.GetAll()
	if err != nil {
		log.Error(err)
		pkg.OnError(ctx, err.Error())
	}

	log.Debugf("found %d users from db", len(*users))
	pkg.OnSuccessWithData(ctx, *users)
}
