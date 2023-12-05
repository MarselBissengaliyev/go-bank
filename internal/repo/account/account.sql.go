package account

import (
	"context"
	"time"

	"github.com/MarselBissengaliyev/bank/internal/models"
	"github.com/jackc/pgx/v5"
)

type AccountSQL struct {
	db *pgx.Conn
}

func NewAccountSQL(db *pgx.Conn) *AccountSQL {
	return &AccountSQL{db}
}

const createAccount = `
	INSERT INTO accounts (
		owner,
		balance,
		currency
	) VALUES (
		$1, $2, $3
	) RETURNING id, owner, balance, currency, created_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (r *AccountSQL) CreateAccount(arg CreateAccountParams) (models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, createAccount, arg.Owner, arg.Balance, arg.Currency)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)

	return i, err
}

const getAccount = `
	SELECT id, owner, balance, currency, created_at FROM accounts 
	WHERE id = $1
`

func (r *AccountSQL) GetAccount(id int64) (models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, getAccount, id)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)

	return i, err
}

const listAccounts = `
	SELECT id, owner, balance, currency, created_at FROM accounts
	ORDER BY id
	LIMIT $1
	OFFSET $2
`

type ListAccountParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (r *AccountSQL) ListAccounts(arg ListAccountParams) ([]models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := r.db.Query(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []models.Account
	for rows.Next() {
		var i models.Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	rows.Close()

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

const updateAccount = `
	UPDATE accounts
	SET balance = $2
	WHERE id = $1
	RETURNING *
`

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

func (r *AccountSQL) UpdateAccount(arg UpdateAccountParams) (models.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	row := r.db.QueryRow(ctx, updateAccount, arg.ID, arg.Balance)
	var i models.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)

	return i, err
}

const deleteAccount = `
	DELETE FROM accounts
	WHERE id = $1
`

func (r *AccountSQL) DeleteAccount(id int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.db.Exec(ctx, deleteAccount, id)
	return err
}
