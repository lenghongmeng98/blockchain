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

	/*
		endpoint: http://localhost:8080/api/v1/stu/create
		body:
		{
		    "stuname": "hongmeng",
		    "gender": "male",
		    "ClassName": "sr class",
		    "Note": "hello"
		}
	*/
	router.POST("/create", controller.StuCreate)

	// endpoint: http://localhost:8080/api/v1/stu/by-id/9066f0c0bbbece4e61e84a0c3d026434
	router.GET("/by-id/:id", controller.StuFindById)

	// endpoint: http://localhost:8080/api/v1/stu/filter?stuname=hong
	// endpoint: http://localhost:8080/api/v1/stu/filter?classname=sr
	router.GET("/filter", controller.StuFilter)

	/*
		endpoint: http://localhost:8080/api/v1/stu/update/9066f0c0bbbece4e61e84a0c3d026434
		body:
		{
			"_rev" : "1-e7baac476b454c0aa5657e36ae05470c",
			"stuname": "value updated",
			"gender": "Male",
			"classname": "SR class",
			"note": "Test"
		}
	*/
	router.PUT("/update/:id", controller.StuUpdate)

	// endpoint: http://localhost:8080/api/v1/stu/delete/9066f0c0bbbece4e61e84a0c3d00402c/1-d3627ff902b6863710c4329f3b4e32a7
	router.DELETE("/delete/:id/:rev", controller.StuDelete)

	router.POST("/file/upload", controller.UploadFile)
	router.GET("/file/get", controller.GetFile)
}
