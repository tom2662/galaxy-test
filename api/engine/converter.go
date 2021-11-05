package engine

import (
	"galaxy/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConvertNum(c *gin.Context) {

	res := "hello world"

	c.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Result set of Post",
		Data:    res})
}
