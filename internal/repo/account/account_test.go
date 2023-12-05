package account_test

import (
	"log"
	"testing"
	"time"

	"github.com/MarselBissengaliyev/bank/config"
	"github.com/MarselBissengaliyev/bank/db/postgres"
	"github.com/MarselBissengaliyev/bank/internal/models"
	"github.com/MarselBissengaliyev/bank/internal/repo"
	"github.com/MarselBissengaliyev/bank/internal/repo/account"
	"github.com/MarselBissengaliyev/bank/util"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/require"
)

var testQueries = func(t *testing.T) *repo.Repository {
	postgresCfg, err := config.InitPostgresConfig("../../../config/")

	if err != nil {
		log.Fatal("Failed to initialize postgres config: ", err)
	}

	conn, err := postgres.NewPostgresDB(postgresCfg)
	if err != nil {
		require.NoError(t, err)
	}

	repo := repo.NewRepository(conn)

	return repo
}

func createRandomAccount(t *testing.T) models.Account {
	arg := account.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	queries := testQueries(t)
	account, err := queries.Account.CreateAccount(arg)
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
	account1 := createRandomAccount(t)
	queries := testQueries(t)

	account2, err := queries.Account.GetAccount(account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	queries := testQueries(t)

	arg := account.UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := queries.Account.UpdateAccount(arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	queries := testQueries(t)

	err := queries.DeleteAccount(account1.ID)
	require.NoError(t, err)

	account2, err := queries.Account.GetAccount(account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, pgx.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	queries := testQueries(t)
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := account.ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := queries.Account.ListAccounts(arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
