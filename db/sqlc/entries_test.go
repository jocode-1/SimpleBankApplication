package db

import (
	"context"
	"github.com/jocode-1/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntries(t, account)
}

func TestGetEntries(t *testing.T) {

	account := createRandomAccount(t)
	entries := createRandomEntries(t, account)
	entries2, err := testQueries.GetEntries(context.Background(), entries.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entries2)

	require.Equal(t, entries.ID, entries2.ID)
	require.Equal(t, entries.AccountID, entries2.AccountID)
	require.Equal(t, entries.Amount, entries2.Amount)
	require.WithinDuration(t, entries.CreatedAt, entries2.CreatedAt, time.Second)

}

func TestUpdateEntries(t *testing.T) {

	account := createRandomAccount(t)
	entries := createRandomEntries(t, account)
	args := UpdateEntriesParams{
		ID:     entries.ID,
		Amount: util.RandomMoney(),
	}
	entries1, err := testQueries.UpdateEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries1)

	require.Equal(t, entries.ID, entries1.ID)
	require.Equal(t, args.Amount, entries1.Amount)
	require.Equal(t, entries.CreatedAt, entries1.CreatedAt, time.Now())

}

func TestListEntries(t *testing.T) {

	for i := 0; i < 10; i++ {
		entries := createRandomAccount(t)
		createRandomEntries(t, entries)
	}
	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries1, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries1, 5)

	for _, entries := range entries1 {
		require.NotEmpty(t, entries)
	}

}
