package user_services

import (
	"api/internal/model"
	"testing"
)

var testData = Db{
	"1": model.Todo{
		Key:   "test",
		Value: "test",
	},
}

func CheckError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func Check(t *testing.T, got model.Todo, want model.Todo) {
	t.Helper()
	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func CheckErrorEqual(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got = %s want = %s", got, want)
	}
}

func TestDb_Search(t *testing.T) {
	got, err := testData.Search("1")
	CheckError(t, err)
	want := model.Todo{Key: "test", Value: "test"}
	Check(t, got, want)
}

func TestDb_Add(t *testing.T) {
	err := testData.Add("2", model.Todo{Key: "testkey", Value: "testValue"})
	CheckError(t, err)
	got, err := testData.Search("2")
	CheckError(t, err)
	want := model.Todo{Key: "testkey", Value: "testValue"}
	Check(t, got, want)
}

func TestDb_Remove(t *testing.T) {
	err := testData.Remove("2")
	CheckError(t, err)
	_, got := testData.Search("2")
	want := UsKeyNotFound
	CheckErrorEqual(t, got, want)
}
