package user_services

import "testing"

var testData = Db{
	"test": "test",
}

func CheckError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Check(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got = %s want = %s", got, want)
	}
}

func TestDb_Search(t *testing.T) {
	got, err := testData.Search("test")
	CheckError(t, err)
	want := "test"
	Check(t, got, want)
}

func TestDb_Add(t *testing.T) {
	err := testData.Add("key", "value")
	CheckError(t, err)
	got, err := testData.Search("key")
	CheckError(t, err)
	want := "value"
	Check(t, got, want)
}

func TestDb_Remove(t *testing.T) {
	err := testData.Remove("key")
	CheckError(t, err)
	_, got := testData.Search("key")
	want := UsKeyNotFound
	Check(t, got.Error(), want.Error())
}
