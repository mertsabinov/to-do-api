package user_service

type UserServicesError string

const (
	UsKeyNotFound      UserServicesError = "This key is not found"
	UsKeyAllReadyExist UserServicesError = "this key is all ready exist"
)

func (us UserServicesError) Error() string {
	return string(us)
}
