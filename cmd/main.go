package main

import (
	"api/internal/model"
	"api/internal/user_controller"
	"api/internal/user_services"
	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	r := gin.Default()
	us := user_service.NewUserService([]model.Todo{})
	uc := user_controller.NewUserConroller(us)
	basepath := r.Group("/v1")
	uc.UserControllerRout(basepath)
	return r
}

func main() {
	r := Start()
	r.Run(":8080")
}
