package repsonse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应结构体
type Response struct {
	Ret     int         `json:"ret"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// Success 成功时响应
func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{0, "success", data})
}

// Fail 失败时响应
func Fail(ret int, message string, c *gin.Context) {
	c.JSON(200, Response{ret, message, nil})
}
