package controller

import (
	"blockchain/conn"
	"blockchain/model"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func StuCreate(c *gin.Context) {

	var stu model.Stu
	c.Bind(&stu)
	id, _, err := conn.DBConn.CreateDoc(context.TODO(), stu)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   FindById(id),
	})
}

func StuFindById(c *gin.Context) {
	id := c.Param("id")

	var stu model.Stu
	err := conn.DBConn.Get(context.TODO(), id).ScanDoc(&stu)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   stu,
	})
}

func FindById(id string) model.Stu {
	fmt.Print(id)
	var stu model.Stu
	err := conn.DBConn.Get(context.TODO(), id).ScanDoc(&stu)

	if err != nil {
		panic(err)
	}

	return stu
}

func StuFindByRevId(c *gin.Context) {
	rev := c.Param("rev")

	var stu model.Stu
	err := conn.DBConn.Get(context.TODO(), rev).ScanDoc(&stu)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   stu,
	})
}

func StuUpdate(c *gin.Context) {

	id := c.Param("id")
	var stu model.Stu
	c.Bind(&stu)
	newRev, err := conn.DBConn.Put(context.TODO(), id, stu)

	fmt.Print(newRev)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status": "success",
		"data":   FindById(id),
	})
}
