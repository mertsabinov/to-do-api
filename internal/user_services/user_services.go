package user_services

import (
	"api/internal/model"
)

type Db map[string]model.Todo

func (db Db) Search(key string) (model.Todo, error) {
	result, ok := db[key]
	if !ok {
		return model.Todo{}, UsKeyNotFound
	}
	return result, nil
}

func (db Db) Add(key string, value model.Todo) error {
	_, err := db.Search(key)
	switch err {
	case nil:
		return UsKeyAllReadyExist
	case UsKeyNotFound:
		db[key] = value
		break
	}
	return nil
}

func (db Db) Remove(key string) error {
	_, err := db.Search(key)
	switch err {
	case nil:
		delete(db, key)
		break
	case UsKeyNotFound:
		return UsKeyNotFound
	}
	return nil
}

func (db Db) GetAll() Db {
	return db
}
