package users

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isAdmin"`
}

func NewUser(name string, email string, password string, isAdmin bool) (*User, error) {
	user := User{
		Name:     name,
		Email:    email,
		Password: password,
		IsAdmin:  isAdmin,
	}

	user.prepare()

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Validate() error {
	if u.ID == "" {
		return fmt.Errorf("user id is required")
	}

	if u.Name == "" {
		return fmt.Errorf("user name is required")
	}

	if u.Email == "" {
		return fmt.Errorf("e-mail is required")
	}

	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, u.Email)
	if !match {
		return fmt.Errorf("e-mail is not valid")
	}

	return nil
}

func (u *User) prepare() {
	u.ID = uuid.NewString()
}
