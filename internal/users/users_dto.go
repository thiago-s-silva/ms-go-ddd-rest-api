package users

import (
	"fmt"
	"regexp"
)

type NewUserRequestDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

type NewUserResponseDTO struct {
	ID string `json:"id"`
}

type ErrorReponseDTO struct {
	Message string `json:"message"`
}

func (dto *NewUserRequestDTO) Validate() error {
	if dto.Name == "" {
		return fmt.Errorf("user name is required")
	}

	if dto.Email == "" {
		return fmt.Errorf("e-mail is required")
	}

	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, dto.Email)
	if !match {
		return fmt.Errorf("e-mail is not valid")
	}

	return nil
}
