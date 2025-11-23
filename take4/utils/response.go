package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger 全局日志实例
var Logger = logrus.New()

func init() {
	// 设置日志格式
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}
