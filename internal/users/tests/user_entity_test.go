package user_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thiago-s-silva/ms-go-task-management-api/internal/users"
)

func TestUserValidation(t *testing.T) {
	var u = users.User{}

	err := u.Validate().Error()
	require.NotNil(t, err)

}

func TestUserValidatonWhenValid(t *testing.T) {
	_, err := users.NewUser("Test", "test@gmail.com", "abc", false)

	require.Nil(t, err)
}
