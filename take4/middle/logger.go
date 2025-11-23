package middle

import (
	"time"

	"homework/take4/utils"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 记录请求日志
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()

		utils.Logger.WithFields(map[string]interface{}{
			"status_code": statusCode,
			"latency":     latency,
			"client_ip":   c.ClientIP(),
			"method":      c.Request.Method,
			"path":        c.Request.URL.Path,
		}).Info("Request processed")
	}
}

// ErrorHandlerMiddleware 错误处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			utils.Logger.WithFields(map[string]interface{}{
				"error": err.Error(),
				"path":  c.Request.URL.Path,
			}).Error("Request error")
		}
	}
}
