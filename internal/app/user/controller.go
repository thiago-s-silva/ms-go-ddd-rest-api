package user

import (
	"github.com/gin-gonic/gin"
	"github.com/thiago-s-silva/ms-go-task-management-api/pkg"
)

func GetOneHandler(ctx *gin.Context) {
	var urlParams UrlParams
	bindUriParams(ctx, &urlParams)

	s := UserService{context: ctx}
	s.GetOne(urlParams.ID)
}

func CreateHandler(ctx *gin.Context) {
	var payload UserRequestPayload
	bindPayload(ctx, &payload)

	if err := payload.Validate(); err != nil {
		pkg.OnBadRequest(ctx, err.Error())
	}

	s := UserService{context: ctx}
	s.Create(&payload)
}

func bindUriParams(ctx *gin.Context, u *UrlParams) {
	if err := ctx.ShouldBindUri(&u); err != nil {
		pkg.OnBadRequest(ctx, err.Error())
	}
}

func bindPayload(ctx *gin.Context, u *UserRequestPayload) {
	if err := ctx.ShouldBind(&u); err != nil {
		pkg.OnBadRequest(ctx, "invalid body request")
	}
}
