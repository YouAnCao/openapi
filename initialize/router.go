package initialize

import (
	"fmt"
	"net/http"
	"openapi/api"
	"openapi/global"
	"openapi/middleware"

	"github.com/gin-gonic/gin"
)

// Router 路由方法
func Router() {
	engine := gin.Default()

	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found!")
	})

	openapi := engine.Group("/api")
	{
		openapi.POST("/login", api.GetUser().UserLogin)
		openapi.POST("/sign", api.GetUser().UserRegister)
		openapi.Use(middleware.JwtAuth())
		openapi.GET("/user", api.GetUser().GetUserInfo)
	}

	engine.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Server.Port))

}
