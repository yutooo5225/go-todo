package db

import (
	"context"
	"database/sql"
	"go-todo/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) Task {
	list := createRandomList(t)

	arg := CreateTaskParams{
		TaskID:      util.RandomId(),
		ListID:      sql.NullInt32{list.ListID, true},
		Description: sql.NullString{util.RandomDescription(), true},
		Done:        sql.NullBool{util.RandomBool(), true},
		CreateAt:    sql.NullTime{util.RandomTime(), true},
		UpdateAt:    sql.NullTime{util.RandomTime(), false},
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.Done, task.Done)

	require.NotEmpty(t, task.TaskID)
	require.NotEmpty(t, task.ListID)

	require.NotZero(t, task.CreateAt)

	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := createRandomTask(t)
	task2, err := testQueries.GetTask(context.Background(), task1.TaskID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.TaskID, task2.TaskID)
	require.Equal(t, task1.ListID, task2.ListID)
	require.Equal(t, task1.Description, task2.Description)
	require.Equal(t, task1.Done, task2.Done)
	require.Equal(t, task1.CreateAt, task2.CreateAt)
	require.Equal(t, task1.UpdateAt, task2.UpdateAt)
}

func TestUpdateTask(t *testing.T) {
	task1 := createRandomTask(t)
	list1 := createRandomList(t)

	arg := UpdateTaskParams{
		TaskID:      task1.TaskID,
		ListID:      sql.NullInt32{list1.ListID, true},
		Description: sql.NullString{util.RandomDescription(), true},
		Done:        sql.NullBool{util.RandomBool(), true},
		UpdateAt:    task1.CreateAt,
	}

	task2, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.TaskID, task2.TaskID)
	require.Equal(t, arg.ListID, task2.ListID)
	require.Equal(t, arg.Description, task2.Description)
	require.Equal(t, arg.Done, task2.Done)
	require.Equal(t, arg.UpdateAt, task2.UpdateAt)
}

func TestDeleteTask(t *testing.T) {
	task1 := createRandomTask(t)

	err := testQueries.DeleteTask(context.Background(), task1.TaskID)
	require.NoError(t, err)

	task2, err := testQueries.GetTask(context.Background(), task1.TaskID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}

func TestListTasks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTask(t)
	}

	arg := ListTasksParams{
		Limit:  5,
		Offset: 5,
	}

	tasks, err := testQueries.ListTasks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tasks, 5)

	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}
