package router

import (
	"github.com/gin-gonic/gin"
	"GinWebBase/api"
	)


// 初始化路由,树型结构
func InitRouter(r *gin.Engine)  {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", api.Ping)
		v1.GET("/user/:id", api.GetUser)
		v1.POST("/user", api.AddUser)
		v1.PUT("/user/:id", api.UpdateUser)
		v1.DELETE("/user/:id", api.DeleteUser)
	}
}
