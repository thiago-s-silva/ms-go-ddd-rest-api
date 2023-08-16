package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

func GetOneHandler(ctx *gin.Context) {
	// init a new logger instance
	logger := pkg.NewLogger("GetOne User")

	// bind URL Params based on struct
	var urlParams UrlParams
	if err := bindUriParams(ctx, &urlParams); err != nil {
		logger.Error(err)
		pkg.OnBadRequest(ctx, err.Error())
		return
	}

	// init UserService injecting needed dependencies
	s := UserService{context: ctx, log: logger}
	// call service method to get one user
	s.GetOne(urlParams.ID)
}

func CreateHandler(ctx *gin.Context) {
	// init new logger instance
	logger := pkg.NewLogger("Create One User")

	// bind request json payload based on struct
	var payload UserRequestPayload
	bindPayload(ctx, &payload)

	// validate payload based on dto
	if err := payload.Validate(); err != nil {
		logger.Error(err.Error())
		pkg.OnBadRequest(ctx, err.Error())
		return
	}

	// init new UserService instance injecting needed dependencies
	s := UserService{context: ctx, log: logger}
	// call create method to create a new user
	s.Create(&payload)
}

func bindUriParams(ctx *gin.Context, u *UrlParams) error {
	// try to bind the Request URI
	if err := ctx.ShouldBindUri(&u); err != nil {
		return err
	}

	return nil
}

func bindPayload(ctx *gin.Context, u *UserRequestPayload) error {
	// try to bind Request Payload
	if err := ctx.ShouldBind(&u); err != nil {
		return err
	}

	return nil
}
