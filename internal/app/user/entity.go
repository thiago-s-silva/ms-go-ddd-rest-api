package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Role     string
	Email    string
	Password string
}
