package main

import (
	"homework/take4/config"
	"homework/take4/routes"
	"homework/take4/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 设置 Gin 模式
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 引擎
	r := gin.New()

	// 设置路由
	routes.SetupRoutes(r)

	// 启动服务器
	utils.Logger.Info("Starting server on port ", config.ServerPort)
	log.Printf("Server is running on http://localhost:%s", config.ServerPort)
	log.Printf("API documentation: http://localhost:%s/api/health", config.ServerPort)

	if err := r.Run(":" + config.ServerPort); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
