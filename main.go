package main

import (
	"github.com/gin-gonic/gin"
	"GinWebBase/router"
	"GinWebBase/utils"
	"GinWebBase/model"
	"GinWebBase/config"
	"fmt"
)

func main() {
	// 初始化配置
	r := gin.Default()

	conf := config.Read()                                // 读取配置文件

	router.InitRouter(r)                                 // 初始化路由
	utils.InitLogrus(conf.Log.LogPath, conf.Log.LogName) // 初始化日志

	DB := model.InitDB(conf) // 初始化数据库模型

	defer DB.Close() // 程序结束关闭数据库链接

	// 设置配置监听地址
	addr := fmt.Sprintf("%s:%s", conf.Server.Addr, conf.Server.Port)
	r.Run(addr) // 开启监听服务
}
