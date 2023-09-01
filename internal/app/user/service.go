package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

type IUserService interface {
	Create(ctx *gin.Context, log *pkg.Logger, payload *UserRequestPayload)
	GetAll(ctx *gin.Context, log *pkg.Logger)
	GetOne(ctx *gin.Context, log *pkg.Logger, id int)
	Delete(ctx *gin.Context, log *pkg.Logger, id int)
	Update(ctx *gin.Context, log *pkg.Logger, payload *User, id int)
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

func (u *UserService) GetOne(ctx *gin.Context, log *pkg.Logger, id int) {
	user, err := repository.GetOneById(id)
	if err != nil {
		log.Debugf("failed when tried to get user by id %d: %s", id, err.Error())
		pkg.OnNotFound(ctx, err.Error())
		return
	}

	pkg.OnSuccessWithData(ctx, &user)
}

func (u *UserService) Delete(ctx *gin.Context, log *pkg.Logger, id int) {
	_, err := repository.GetOneById(id)
	if err != nil {
		log.Debugf("failed when tried to get user by id %d: %s", id, err.Error())
		pkg.OnNotFound(ctx, err.Error())
		return
	}

	if err = repository.DeleteByID(id); err != nil {
		log.Errorf("failed when tried to delete user by id %d: %s", id, err.Error())
		pkg.OnError(ctx, err.Error())
		return
	}

	pkg.OnSuccess(ctx, "Successfuly deleted user!")
}

func (u *UserService) Update(ctx *gin.Context, log *pkg.Logger, p *UserRequestPayload, id int) {
	_, err := repository.GetOneById(id)
	if err != nil {
		log.Debugf("failed when tried to get user by id %d: %s", id, err.Error())
		pkg.OnNotFound(ctx, err.Error())
		return
	}

	user := User{
		ID:       uint(id),
		Name:     p.Name,
		Role:     p.Role,
		Email:    p.Email,
		Password: p.Password,
	}

	if err = repository.Update(&user); err != nil {
		log.Errorf("failed when tried to update user by id %d: %s", id, err.Error())
		pkg.OnError(ctx, err.Error())
	}

	pkg.OnSuccess(ctx, "successfuly updated user!!")
}
