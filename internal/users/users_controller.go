package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
	The controller is responsible to handle the request, including
*/

type Users interface {
	CreateNewUser(u *NewUserRequestDTO)
}

func CreateNewUser(c echo.Context) error {
	var payload NewUserRequestDTO

	err := c.Bind(&payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorReponseDTO{Message: err.Error()})
	}

	err = payload.Validate()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return createNewUserService(c, &payload)
}
