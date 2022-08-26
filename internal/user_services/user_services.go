package user_services

import (
	"api/internal/model"
	"api/internal/model/request_model"
)

type Db []model.Todo

func (db *Db) Search(id string) (model.Todo, error) {
	var result model.Todo
	for _, todo := range *db {
		if todo.Id == id {
			result = todo
			break
		}
	}
	err := db.CheckNill(result.Id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (db *Db) Add(id string, todo request_model.RequestTodo) error {
	newTodo := model.Todo{
		Id:    id,
		Key:   todo.Key,
		Value: todo.Value,
	}
	_, err := db.Search(id)
	switch err {
	case nil:
		return UsKeyAllReadyExist
	case UsKeyNotFound:
		*db = append(*db, newTodo)
		break
	}
	return nil
}

func (db *Db) Remove(id string) error {
	_, err := db.Search(id)
	newDb := []model.Todo{}
	switch err {
	case nil:
		for _, todo := range *db {
			if todo.Id != id {
				newDb = append(newDb, todo)
			}
		}
		*db = newDb
		break
	case UsKeyNotFound:
		return UsKeyNotFound
	}
	return nil
}

func (db *Db) GetAll() Db {
	return *db
}

func (db Db) CheckNill(id string) error {
	if id == "" {
		return UsKeyNotFound
	}
	return nil
}
