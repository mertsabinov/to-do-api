package user_controller

import (
	"api/internal/model"
	"api/internal/model/request_model"
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
	err := ctx.ShouldBindJSON(&todo)
	newId := uuid.NewString()
	err = us.UserServices.Add(newId, todo)
	result, err := us.UserServices.Search(newId)
	us.CheckError(ctx, err)
	if err == nil {
		ctx.JSONP(http.StatusOK, gin.H{"key": result.Key, "value": result.Value})
	}
}

func (us *UserController) GetAll(ctx *gin.Context) {
	data := us.UserServices.GetAll()
	ctx.JSONP(http.StatusOK, data)
}

func (us *UserController) Delete(ctx *gin.Context) {
	var body request_model.RequestId
	err := ctx.ShouldBindJSON(&body)
	_, err = us.UserServices.Search(body.Id)
	err = us.UserServices.Remove(body.Id)
	us.CheckError(ctx, err)
	if err == nil {
		ctx.JSONP(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (us *UserController) CheckError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSONP(http.StatusBadRequest, gin.H{"message": err})
	}
}

func (us *UserController) UserControllerRout(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.GET("/ping", us.Ping)
	userRoute.POST("/add", us.Add)
	userRoute.GET("/todo", us.GetAll)
	userRoute.DELETE("/todo", us.Delete)
}
