package config

import (
	"os"
)

// JWTSecret JWT 密钥，生产环境应从环境变量读取
var JWTSecret = []byte(getEnv("JWT_SECRET", "your-secret-key-change-in-production"))

// ServerPort 服务器端口
var ServerPort = getEnv("SERVER_PORT", "8080")

// getEnv 获取环境变量
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
