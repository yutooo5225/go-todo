package db

import (
	"context"
	"database/sql"
	"go-todo/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomList(t *testing.T) List {
	arg := CreateListParams{
		ListID:   util.RandomId(),
		ListName: sql.NullString{util.RandomString(4), true},
	}

	job, err := testQueries.CreateList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, job)

	require.Equal(t, arg.ListName, job.ListName)

	require.NotEmpty(t, job.ListID)

	return job
}

func TestCreateList(t *testing.T) {
	createRandomList(t)
}

func TestGetList(t *testing.T) {
	list1 := createRandomList(t)
	list2, err := testQueries.GetList(context.Background(), list1.ListID)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.ListID, list2.ListID)
	require.Equal(t, list1.ListName, list2.ListName)
}

func TestUpdateList(t *testing.T) {
	list1 := createRandomList(t)

	arg := UpdateListParams{
		ListID:   list1.ListID,
		ListName: sql.NullString{util.RandomString(4), true},
	}

	list2, err := testQueries.UpdateList(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.ListID, list2.ListID)
	require.Equal(t, arg.ListName, list2.ListName)
}

func TestDeleteList(t *testing.T) {
	list1 := createRandomList(t)

	err := testQueries.DeleteList(context.Background(), list1.ListID)
	require.NoError(t, err)

	list2, err := testQueries.GetList(context.Background(), list1.ListID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, list2)
}

func TestLists(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomList(t)
	}

	arg := ListsParams{
		Limit:  5,
		Offset: 5,
	}

	lists, err := testQueries.Lists(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, lists, 5)

	for _, list := range lists {
		require.NotEmpty(t, list)
	}
}
