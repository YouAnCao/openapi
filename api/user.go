package api

import (
	"fmt"
	"openapi/modules/user"
	repsonse "openapi/response"

	"openapi/service"

	"github.com/gin-gonic/gin"
)

// User 用户实体
type User struct {
}

// GetUser 获取对象
func GetUser() *User {
	return &User{}
}

// UserLogin 用户登录
func (u *User) UserLogin(c *gin.Context) {
	var loginParm user.LoginParam
	if err := c.ShouldBind(&loginParm); err != nil {
		fmt.Println(err)
		repsonse.Fail(1, "登录失败，请求参数不正确", c)
		return
	}
	userInfo := service.GetUserService().UserLogin(&loginParm)
	if userInfo != nil {
		repsonse.Success(userInfo, c)
		return
	}
	repsonse.Fail(-1, "登录失败，用户名或密码不正确！", c)
}

// UserRegister 用户注册
func (u *User) UserRegister(c *gin.Context) {
	var signParam user.SignParam
	if err := c.ShouldBind(&signParam); err != nil {
		fmt.Println(err)
		repsonse.Fail(1, "参数不正确", c)
		return
	}
	userInfo := service.GetUserService().UserRegister(&signParam)
	if userInfo != nil {
		repsonse.Success(userInfo, c)
		return
	}
	repsonse.Fail(-1, "注册失败!", c)
}
