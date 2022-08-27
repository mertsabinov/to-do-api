package user_service

import (
	"api/internal/model"
	"api/internal/model/model_request"
)

type UserService struct {
	Db []model.Todo
}

func NewUserService(db []model.Todo) UserService {
	return UserService{Db: db}
}

func (us *UserService) Search(id string) (model.Todo, error) {
	var result model.Todo
	for _, todo := range us.Db {
		if todo.Id == id {
			result = todo
			break
		}
	}
	err := us.CheckNill(result.Id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (us *UserService) Add(id string, todo model_request.RequestTodo) error {
	newTodo := model.Todo{
		Id:    id,
		Key:   todo.Key,
		Value: todo.Value,
	}
	_, err := us.Search(id)
	switch err {
	case nil:
		return UsKeyAllReadyExist
	case UsKeyNotFound:
		us.Db = append(us.Db, newTodo)
		break
	}
	return nil
}

func (us *UserService) Remove(id string) error {
	_, err := us.Search(id)
	newDb := []model.Todo{}
	switch err {
	case nil:
		for _, todo := range us.Db {
			if todo.Id != id {
				newDb = append(newDb, todo)
			}
		}
		us.Db = newDb
		break
	case UsKeyNotFound:
		return UsKeyNotFound
	}
	return nil
}

func (us *UserService) GetAll() []model.Todo {
	return us.Db
}

func (db *UserService) CheckNill(id string) error {
	if id == "" {
		return UsKeyNotFound
	}
	return nil
}
