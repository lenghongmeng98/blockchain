package controller

import (
	"blockchain/conn"
	"blockchain/model"
	"blockchain/request"
	"blockchain/response"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-kivik/kivik/v3"
	"github.com/google/uuid"
)

var body request.StuRequest

func StuCreate(c *gin.Context) {

	c.Bind(&body)
	id, _, err := conn.DBConn.CreateDoc(context.TODO(), body)

	if err != nil {
		panic(err)
	}

	response.ResponseBody(c, FindById(id))
}

func StuFindById(c *gin.Context) {
	id := c.Param("id")

	var stu model.Stu
	err := conn.DBConn.Get(context.TODO(), id).ScanDoc(&stu)

	if err != nil {
		panic(err)
	}

	response.ResponseBody(c, stu)
}

func FindById(id string) model.Stu {

	var stu model.Stu
	err := conn.DBConn.Get(context.TODO(), id).ScanDoc(&stu)

	if err != nil {
		panic(err)
	}

	return stu
}

func StuFilter(c *gin.Context) {
	var key string
	var _design string
	var _view string

	var stuname = c.Query("stuname")
	var classname = c.Query("classname")

	if stuname != "" {
		key = stuname
		_design = "_design/by-stu-name"
		_view = "_view/stu-name"
	} else if classname != "" {
		key = classname
		_design = "_design/by-classname"
		_view = "_view/classname"
	} else {
		key = ""
		_design = "_design/by-stu-name"
		_view = "_view/stu-name"
	}

	rows, err := conn.DBConn.Query(context.TODO(), _design, _view, kivik.Options{
		"include_docs": true,
		"startkey":     key,
		"endkey":       key + kivik.EndKeySuffix,
	})

	if err != nil {
		panic(err)
	}

	var students = []model.Stu{}

	for rows.Next() {

		var stu model.Stu

		if err := rows.ScanDoc(&stu); err != nil {
			panic(err)
		}

		students = append(students, stu)

	}

	if rows.Err() != nil {
		panic(rows.Err())
	}

	response.ResponseBodyV2(c, students)
}

func StuUpdate(c *gin.Context) {

	id := c.Param("id")
	c.Bind(&body)
	newRev, err := conn.DBConn.Put(context.TODO(), id, body)

	fmt.Print(newRev)
	if err != nil {
		panic(err)
	}

	response.ResponseBody(c, FindById(id))
}

func StuDelete(c *gin.Context) {

	id := c.Param("id")
	rev := c.Param("rev")

	newRev, err := conn.DBConn.Delete(context.TODO(), id, rev)

	fmt.Print(newRev)
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "deleted",
	})
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")

	conn.DBConn.PutAttachment(context.TODO(), uuid.NewString(), file.Filename, "multipart/form-data")

}
