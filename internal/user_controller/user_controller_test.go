package user_controller

import (
	"api/internal/user_services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var Url string = "http://localhost:8080/"

func StartTest() *gin.Engine {
	r := gin.Default()
	us := user_services.Db{}
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

func ConvertMapToJsonString(t *testing.T, want map[string]string) string {
	wantByte, err := json.Marshal(want)
	CheckError(t, err)
	return string(wantByte)
}

func TestUserController_Ping(t *testing.T) {
	router := StartTest()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", Url+"v1/user/ping", nil)
	router.ServeHTTP(w, req)
	want := map[string]string{
		"message": "pong",
	}
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, ConvertMapToJsonString(t, want), w.Body.String())
}
