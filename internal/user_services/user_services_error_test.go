package user_services

import "testing"

func CheckUserServicesError(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got = %s want = %s", got, want)
	}
}

func TestUserServicesError_Error(t *testing.T) {
	t.Run("UsKeyNotFound", func(t *testing.T) {
		got := UsKeyNotFound.Error()
		want := "This key is not found"
		CheckUserServicesError(t, got, want)
	})
	t.Run("UsKeyAllReadyExist", func(t *testing.T) {
		got := UsKeyAllReadyExist.Error()
		want := "this key is all ready exist"
		CheckUserServicesError(t, got, want)
	})
}
