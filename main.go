package main

import (
	"blockchain/conn"
	"blockchain/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	conn.ConnectToDB()
}

func main() {
	r := gin.Default()
	RegisterRoutes(r)
	r.Run()
}

func RegisterRoutes(r *gin.Engine) {
	router := r.Group("api/v1/stu/")
	router.POST("/create", controller.StuCreate)
	router.GET("/by-id/:id", controller.StuFindById)
	router.GET("/by-rev/:rev", controller.StuFindByRevId)
	router.PUT("/update/:id", controller.StuUpdate)
}
