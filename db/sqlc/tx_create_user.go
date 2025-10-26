package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

// Transaction for CreateUsering money between accounts
func (store *SQLStore) CreateUserTx(ctx context.Context, args CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.executeTransaction(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, args.CreateUserParams)
		if err != nil {
			return err
		}

		return args.AfterCreate(result.User)
	})

	return result, err
}
