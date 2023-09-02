package controller

import (
	"blockchain/conn"
	"blockchain/model"
	"blockchain/request"
	"blockchain/response"
	"context"
	"fmt"
	"net/http"

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

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from form"})
		panic(err)
	}

	// Open the uploaded file
	uploadedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		panic(err)
	}
	defer uploadedFile.Close()

	att := &kivik.Attachment{
		Content:     uploadedFile,
		Filename:    file.Filename,
		ContentType: file.Header.Get("Content-Type"),
	}

	doc := conn.DBConn.Get(context.Background(), "document_id", kivik.Options{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rev id"})
		panic(err)
	}
	revID := doc.Rev

	// Put the attachment into the document
	_, err = conn.DBConn.PutAttachment(context.Background(), uuid.NewString(), revID, att)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to put attachment"})
		panic(err)
	}

	fmt.Print(att)

	c.JSON(http.StatusOK, gin.H{"message": "Attachment uploaded successfully"})
}

func GetFile(c *gin.Context) {

	id := c.Param("id")
	filename := c.Param("filename")

	// Fetch the attachment
	attachment, err := conn.DBConn.GetAttachment(context.Background(), id, filename, kivik.Options{})
	if err != nil {
		fmt.Printf("Failed to get attachment: %v\n", err)
		return
	}

	fmt.Println(attachment)
}
