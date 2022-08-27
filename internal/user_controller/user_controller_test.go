package user_controller

import (
	"api/internal/model"
	"api/internal/model/model_request"
	user_service "api/internal/user_services"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var Url string = "http://localhost:8080/"
var Router = StartTest()

func StartTest() *gin.Engine {
	r := gin.Default()
	us := user_service.NewUserService([]model.Todo{{
		Id:    "1",
		Key:   "test key",
		Value: "test value",
	}})
	uc := NewUserConroller(us)

	basepath := r.Group("/v1")
	uc.UserControllerRout(basepath)
	return r
}

func CheckError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func ConvertStructToJson(t *testing.T, want model_request.RequestTodo) string {
	wantByte, err := json.Marshal(want)
	CheckError(t, err)
	return string(wantByte)
}

func ConvertMapToJsonString(t *testing.T, want map[string]string) string {
	wantByte, err := json.Marshal(want)
	CheckError(t, err)
	return string(wantByte)
}

func ConvertTodoToIoReader(t *testing.T, want model_request.RequestTodo) *bytes.Reader {
	requestByte, _ := json.Marshal(want)
	requestReader := bytes.NewReader(requestByte)
	return requestReader
}

func ConvertTestRequestIdToIoReader(t *testing.T, want model_request.RequestId) *bytes.Reader {
	requestByte, _ := json.Marshal(want)
	requestReader := bytes.NewReader(requestByte)
	return requestReader
}

func TestUserController_Ping(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", Url+"v1/user/ping", nil)
	Router.ServeHTTP(w, req)
	want := map[string]string{
		"message": "pong",
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ConvertMapToJsonString(t, want), w.Body.String())
}

func TestUserController_GetAll(t *testing.T) {
	var got user_service.UserService
	want := user_service.UserService{
		Db: []model.Todo{
			{
				Id:    "1",
				Key:   "test key",
				Value: "test value",
			},
		},
	}
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", Url+"v1/user/todo", nil)
	CheckError(t, err)
	Router.ServeHTTP(w, req)
	err = json.Unmarshal(w.Body.Bytes(), &got)
	CheckError(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func TestUserController_Add(t *testing.T) {
	want := model_request.RequestTodo{
		Key:   "testKey",
		Value: "testValue",
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", Url+"v1/user/todo", ConvertTodoToIoReader(t, want))
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ConvertStructToJson(t, want), w.Body.String())
}

func TestUserController_Delete(t *testing.T) {
	t.Run("Delete (ok)", func(t *testing.T) {
		testId := model_request.RequestId{
			Id: "1",
		}
		w := httptest.NewRecorder()
		req, err := http.NewRequest("DELETE", Url+"v1/user/todo", ConvertTestRequestIdToIoReader(t, testId))
		CheckError(t, err)
		Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		got := `{"message":"ok"}`
		assert.Equal(t, got, w.Body.String())
	})

	t.Run("Delete (Error)", func(t *testing.T) {
		testId := model_request.RequestId{
			Id: "2",
		}
		w := httptest.NewRecorder()
		req, err := http.NewRequest("DELETE", Url+"v1/user/todo", ConvertTestRequestIdToIoReader(t, testId))
		CheckError(t, err)
		Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		got := `{"message":"This key is not found"}`
		assert.Equal(t, got, w.Body.String())
	})
}
