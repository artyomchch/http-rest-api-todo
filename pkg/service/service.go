package service

import (
	http_rest_api_test "github.com/artyomchch/http-rest-api-test"
	"github.com/artyomchch/http-rest-api-test/pkg/repository"
)

type Authorization interface {
	CreateUser(user http_rest_api_test.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list http_rest_api_test.TodoList) (int, error)
	GetAll(userId int) ([]http_rest_api_test.TodoList, error)
	GetById(userId, listId int) (http_rest_api_test.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input http_rest_api_test.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item http_rest_api_test.TodoItem) (int, error)
	GetAll(userId, listId int) ([]http_rest_api_test.TodoItem, error)
	GetById(userId, itemId int) (http_rest_api_test.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input http_rest_api_test.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
