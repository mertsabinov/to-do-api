package user_services

type Db map[string]string

func (db Db) Search(key string) (string, error) {
	result, ok := db[key]
	if !ok {
		return "", UsKeyNotFound
	}
	return result, nil
}

func (db Db) Add(key string, value string) error {
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
