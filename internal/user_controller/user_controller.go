package user_controller

import (
	"api/internal/model"
	"api/internal/user_services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserController struct {
	UserServices user_services.Db
}

func NewUserConroller(UserServices user_services.Db) UserController {
	return UserController{
		UserServices: UserServices,
	}
}

func (us *UserController) Ping(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{"message": "pong"})
}

func (us *UserController) Add(ctx *gin.Context) {
	var todo model.Todo
	ctx.ShouldBindJSON(&todo)
	newId := uuid.NewString()
	us.UserServices.Add(newId, todo)
	result, err := us.UserServices.Search(newId)
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSONP(http.StatusOK, gin.H{"key": result.Key, "value": result.Value})
}

func (us *UserController) GetAll(ctx *gin.Context) {
	data, err := us.UserServices.GetAll()
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSONP(http.StatusOK, data)
}

func (us *UserController) UserControllerRout(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.GET("/ping", us.Ping)
	userRoute.POST("/add", us.Add)
	userRoute.GET("/todo", us.GetAll)
}
