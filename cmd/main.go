package main

import (
	"api/internal/model"
	"api/internal/user_controller"
	"api/internal/user_service"
	"github.com/gin-gonic/gin"
)

func Start() *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	})
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
