package user

import (
	"fmt"
	"regexp"
)

type UrlParams struct {
	ID int `uri:"id" binding:"required"`
}

type UserRequestPayload struct {
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequestPayload) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("property %s [string] is required", "name")
	}

	if u.Role == "" {
		return fmt.Errorf("property %s [string] is required", "role")
	}

	if u.Email == "" {
		return fmt.Errorf("property %s [string] is required", "email")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	if u.Email == "" || !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("property %s [string] is invalid", "email")
	}

	if u.Password == "" {
		return fmt.Errorf("property %s [string] is required", "password")
	}

	return nil
}

type UserErrorResponse struct {
	Message string `json:"message"`
}

type UserSuccessResponse struct {
	Message string `json:"message"`
}

type UserSuccessResponseWithData struct {
	Data UserRequestPayload `json:"data"`
}

type UsersSuccessResponseWithData struct {
	Data []User `json:"data"`
}
