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
	}
}
