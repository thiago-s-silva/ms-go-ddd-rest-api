package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

func DeleteHandler(ctx *gin.Context) {
	// init a new logger instance
	logger := pkg.NewLogger("Delete user by id")

	// bind URL Params based on struct
	var urlParams UrlParams
	if err := bindUriParams(ctx, &urlParams); err != nil {
		logger.Error(err)
		pkg.OnBadRequest(ctx, err.Error())
		return
	}

	// call User Service delete method
	service.Delete(ctx, logger, urlParams.ID)
}

func GetOneHandler(ctx *gin.Context) {
	// init a new logger instance
	logger := pkg.NewLogger("Get user by id")

	// bind URL Params based on struct
	var urlParams UrlParams
	if err := bindUriParams(ctx, &urlParams); err != nil {
		logger.Error(err)
		pkg.OnBadRequest(ctx, err.Error())
		return
	}

	// call User Service get one method
	service.GetOne(ctx, logger, urlParams.ID)
}

func GetAllHandler(ctx *gin.Context) {
	// init new logger instance
	logger := pkg.NewLogger("Get all users")

	// cal User Service get all method
	service.GetAll(ctx, logger)
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

	// call User Service create method
	service.Create(ctx, logger, &payload)
}

func UpdateHandler(ctx *gin.Context) {
	// init a new logger instance
	logger := pkg.NewLogger("Update user by id")

	// bind URL Params based on struct
	var urlParams UrlParams
	if err := bindUriParams(ctx, &urlParams); err != nil {
		logger.Error(err)
		pkg.OnBadRequest(ctx, err.Error())
		return
	}

	// bind request json payload based on struct
	var payload UserRequestPayload
	bindPayload(ctx, &payload)

	// validate payload based on dto
	if err := payload.Validate(); err != nil {
		logger.Error(err.Error())
		pkg.OnBadRequest(ctx, err.Error())
		return
	}

	// call User Service create method
	service.Update(ctx, logger, &payload, urlParams.ID)
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
