package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) executeTransaction(ctx context.Context, fn func(*Queries) error) error {
	transaction, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(transaction)
	err = fn(q)
	if err != nil {
		if txErr := transaction.Rollback(); txErr != nil {
			return fmt.Errorf("Transaction error: %v, Rollback error: %v", err, txErr)
		}
		return err
	}
	return transaction.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Ammount       int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `jsno:"to_entry"`
}

// Transaction for transfering money between accounts
func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.executeTransaction(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Ammount:       args.Ammount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountID,
			Ammount:   -args.Ammount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToAccountID,
			Ammount:   args.Ammount,
		})
		if err != nil {
			return err
		}

		//get account -> update its balance (needs locking)
		if args.FromAccountID < args.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, args.FromAccountID, -args.Ammount, args.ToAccountID, args.Ammount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, args.ToAccountID, args.Ammount, args.FromAccountID, -args.Ammount)
		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:      accountID1,
		Ammount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:      accountID2,
		Ammount: amount2,
	})

	return
}
