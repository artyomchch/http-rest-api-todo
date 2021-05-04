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
	Delete(userId, listId int) error
	Update(userId, listId int, input http_rest_api_test.UpdateListInput) error
}

type TodoItem interface {
	Create(ListId int, item http_rest_api_test.TodoItem) (int, error)
	GetAll(userId, listId int) ([]http_rest_api_test.TodoItem, error)
	GetById(userId, itemId int) (http_rest_api_test.TodoItem, error)
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
