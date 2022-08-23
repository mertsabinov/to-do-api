package user_controller

import (
	"api/internal/user_services"
	"github.com/gin-gonic/gin"
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

func (us *UserController) UserControllerRout(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.GET("/ping", us.Ping)
}
