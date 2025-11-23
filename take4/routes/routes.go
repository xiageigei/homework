package routes

import (
	"homework/take4/controlles"
	"homework/take4/middle"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine) {
	// 应用全局中间件
	r.Use(middle.LoggerMiddleware())
	r.Use(middle.ErrorHandlerMiddleware())
	r.Use(gin.Recovery())

	// API 路由组
	api := r.Group("/api")
	{
		// 健康检查
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok"})
		})

		// 认证相关路由（不需要登录）
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// 需要认证的路由
		authorized := api.Group("")
		authorized.Use(middle.AuthMiddleware())
		{
			// 用户相关
			authorized.GET("/profile", controllers.GetProfile)

			// 文章相关
			authorized.POST("/posts", controllers.CreatePost)
			authorized.PUT("/posts/:id", controllers.UpdatePost)
			authorized.DELETE("/posts/:id", controllers.DeletePost)

			// 评论相关
			authorized.POST("/comments", controllers.CreateComment)
		}

		// 公开访问的路由
		public := api.Group("")
		{
			// 文章列表和详情（不需要登录即可查看）
			public.GET("/posts", controllers.GetPosts)
			public.GET("/posts/:id", controllers.GetPost)

			// 评论相关路由
			comments := public.Group("/comments")
			{
				// 获取文章的评论列表（不需要登录即可查看）
				comments.GET("/post/:post_id", controllers.GetCommentsByPost)
			}
		}
	}
}
