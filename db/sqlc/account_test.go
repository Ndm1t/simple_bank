package db

import (
	"bankingApp/util"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	res, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.ID, account.ID)
	require.Equal(t, res.Owner, account.Owner)
	require.Equal(t, res.Balance, account.Balance)
	require.Equal(t, res.Currency, account.Currency)
	require.Equal(t, res.CreatedAt, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	res, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.ID, account.ID)
	require.Equal(t, res.Balance, arg.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	res, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, res)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
