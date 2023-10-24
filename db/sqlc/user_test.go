package db

import (
	"bankingApp/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(5))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(), //randomly generated
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.Zero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	res, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.Username, user.Username)
	require.Equal(t, res.HashedPassword, user.HashedPassword)
	require.Equal(t, res.FullName, user.FullName)
	require.Equal(t, res.PasswordChangedAt, user.PasswordChangedAt)
	require.Equal(t, res.CreatedAt, user.CreatedAt)
}
