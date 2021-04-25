package repository

import (
	http_rest_api_test "github.com/artyomchch/http-rest-api-test"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user http_rest_api_test.User) (int, error)
	GetUser(username, password string) (http_rest_api_test.User, error)
}

type TodoList interface {
	Create(userId int, list http_rest_api_test.TodoList) (int, error)
	GetAll(userId int) ([]http_rest_api_test.TodoList, error)
	GetById(userId, listId int) (http_rest_api_test.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
