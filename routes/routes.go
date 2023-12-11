package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/users"
)

func RegisterRoutes(e *echo.Echo) {
	// Users endpoints
	e.POST("api/v1/users", users.CreateNewUser)
}
