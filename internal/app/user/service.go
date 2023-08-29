package user

import (
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

type IUserService interface {
	Create(payload *UserRequestPayload)
	GetOne(id string)
}

type UserService struct{}

func (u *UserService) Create(ctx *gin.Context, log *pkg.Logger, p *UserRequestPayload) {
	newUserId, err := uuid.NewRandom()
	if err != nil {
		log.Errorf("failed to generate a new uuid: %s", err.Error())
		pkg.OnError(ctx, err.Error())
	}

	user := User{
		ID:       uint(newUserId.ID()),
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
