package response

import (
	"blockchain/model"
	"time"

	"github.com/gin-gonic/gin"
)

func ResponseBody(c *gin.Context, data model.Stu) {
	c.JSON(200, gin.H{
		"payload":   data,
		"timestamp": time.Now(),
		"status":    "success",
	})
}

func ResponseBodyV2(c *gin.Context, data []model.Stu) {
	c.JSON(200, gin.H{
		"payload":   data,
		"timestamp": time.Now(),
		"status":    "success",
	})
}
