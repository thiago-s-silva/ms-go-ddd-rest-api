package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiago-s-silva/ms-go-task-management-api/utils"
)

func createNewUserService(c echo.Context, dto *NewUserRequestDTO) error {
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorReponseDTO{Message: "error when trying to hash the password"})
	}

	newUser, _ := NewUser(dto.Name, dto.Email, string(hashedPassword), dto.IsAdmin)

	response := NewUserResponseDTO{ID: newUser.ID}
	return c.JSON(http.StatusCreated, response)
}
