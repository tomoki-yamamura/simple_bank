package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 10
	amount := int64(10)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i % 2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID: toAccountID,
				Amount: amount,
			})
			errs <- err
			// results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		// result := <-results
		// require.NotEmpty(t, result)

		// transfer := result.Transfer
		// require.NotEmpty(t, transfer)
		// require.Equal(t, account1.ID, transfer.FromAccountID)
		// require.Equal(t, account2.ID, transfer.ToAccountID)
		// require.Equal(t, amount, transfer.Amount)
		// require.NotZero(t,transfer.ID)
		// require.NotZero(t, transfer.CreatedAt)

		// _, err = store.GetTransfer(context.Background(), transfer.ID)
		// require.NoError(t, err)

		// FromEntry := result.FromEntry
		// require.NotEmpty(t, FromEntry)
		// require.Equal(t, account1.ID, FromEntry.AccountID)
		// require.Equal(t, -amount, FromEntry.Amount)
		// require.NotZero(t, FromEntry.ID)
		// require.NotZero(t, FromEntry.CreatedAt)

		// _, err = store.GetEntry(context.Background(), FromEntry.ID)
		// require.NoError(t, err)

		// ToEntry := result.ToEntry
		// require.NotEmpty(t, ToEntry)
		// require.Equal(t, account2.ID, ToEntry.AccountID)
		// require.Equal(t, amount, ToEntry.Amount)
		// require.NotZero(t, ToEntry.ID)
		// require.NotZero(t, ToEntry.CreatedAt)

		// _, err = store.GetEntry(context.Background(), ToEntry.ID)
		// require.NoError(t, err)

		// fromAccount := result.FromAccount
		// require.NotEmpty(t, fromAccount)
		// require.Equal(t, account1.ID, fromAccount.ID)

		// toAccount := result.ToAccount
		// require.NotEmpty(t, toAccount)
		// require.Equal(t, account2.ID, toAccount.ID)

		// diff1 := account1.Balance - fromAccount.Balance
		// diff2 := toAccount.Balance - account2.Balance
		// require.Equal(t, diff1, diff2)
		// require.True(t, diff1 > 0)
		// require.True(t, diff1 % amount == 0)
		// k := int(diff1 / amount)
		// require.True(t, k >= 1 && k <= n)
		// require.NotContains(t, existed, k)
		// existed[k] = true
	}
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}