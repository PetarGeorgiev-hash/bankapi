package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hasedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hasedPassword)

	err = CheckPassword(password, hasedPassword)
	require.NoError(t, err)

	wrongPassword := RandomString(7)
	err = CheckPassword(wrongPassword, hasedPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hasedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hasedPassword2)
	require.NotEqual(t, hasedPassword, hasedPassword2)

}
