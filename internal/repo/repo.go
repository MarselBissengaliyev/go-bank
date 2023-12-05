package repo

import (
	"github.com/MarselBissengaliyev/bank/internal/models"
	"github.com/MarselBissengaliyev/bank/internal/repo/account"
	"github.com/jackc/pgx/v5"
)

type Account interface {
	CreateAccount(arg account.CreateAccountParams) (models.Account, error)
	GetAccount(id int64) (models.Account, error)
	ListAccounts(arg account.ListAccountParams) ([]models.Account, error)
	UpdateAccount(arg account.UpdateAccountParams) (models.Account, error)
	DeleteAccount(id int64) error
}

type Repository struct {
	Account
}

func NewRepository(db *pgx.Conn) *Repository {
	return &Repository{
		Account: account.NewAccountSQL(db),
	}
}
