package middleware

import (
	"openapi/common"
	repsonse "openapi/response"

	"github.com/gin-gonic/gin"
)

// JwtAuth jwt认证
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			repsonse.Fail(1, "用户未登录", c)
			c.Abort()
			return
		}
		username, err := common.VerifyToken(token)
		if err != nil {
			repsonse.Fail(2, "用户未登录", c)
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
