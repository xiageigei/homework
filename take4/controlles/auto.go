package controllers

import (
	"net/http"

	"homework/take4/config"
	"homework/take4/models"
	"homework/take4/utils"

	"github.com/gin-gonic/gin"
)

// RegisterRequest 注册请求结构
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse 认证响应结构
type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Register 用户注册
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	if err := config.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "用户名已存在")
		return
	}

	// 检查邮箱是否已存在
	if err := config.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "邮箱已被使用")
		return
	}

	// 创建新用户
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
	}

	// 加密密码
	if err := user.HashPassword(req.Password); err != nil {
		utils.Logger.Error("密码加密失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "服务器内部错误")
		return
	}

	// 保存到数据库
	if err := config.DB.Create(&user).Error; err != nil {
		utils.Logger.Error("创建用户失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建用户失败")
		return
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.Logger.Error("生成 token 失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成令牌失败")
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "注册成功", AuthResponse{
		Token: token,
		User:  user,
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	// 查找用户
	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 验证密码
	if !user.CheckPassword(req.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "用户名或密码错误")
		return
	}

	// 生成 JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		utils.Logger.Error("生成 token 失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "生成令牌失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "登录成功", AuthResponse{
		Token: token,
		User:  user,
	})
}

// GetProfile 获取当前用户信息
func GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "未认证")
		return
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "用户不存在")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "获取成功", user)
}
