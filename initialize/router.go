package initialize

import (
	"fmt"
	"net/http"
	"openapi/api"
	"openapi/global"

	"github.com/gin-gonic/gin"
)

// Router 路由方法
func Router() {
	engine := gin.Default()

	engine.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "not found!")
	})

	user := engine.Group("/api")
	{
		user.POST("/login", api.GetUser().UserLogin)
		user.POST("/sign", api.GetUser().UserRegister)
	}

	engine.Run(fmt.Sprintf("0.0.0.0:%d", global.Config.Server.Port))

}
