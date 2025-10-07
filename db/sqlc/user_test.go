package db

import (
	"context"
	"testing"
	"time"

	"github.com/PetarGeorgiev-hash/bankapi/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "password",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}
func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user2.Username, user1.Username)
	require.Equal(t, user2.HashedPassword, user1.HashedPassword)
	require.Equal(t, user2.Email, user1.Email)
	require.WithinDuration(t, user2.PasswordChangedAt, user1.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user2.CreatedAt, user1.CreatedAt, time.Second)

}
