package controllers

import (
	"net/http"
	"strconv"

	"homework/take4/config"
	"homework/take4/models"
	"homework/take4/utils"

	"github.com/gin-gonic/gin"
)

// CreatePostRequest 创建文章请求结构
type CreatePostRequest struct {
	Title   string `json:"title" binding:"required,min=1,max=200"`
	Content string `json:"content" binding:"required,min=1"`
}

// UpdatePostRequest 更新文章请求结构
type UpdatePostRequest struct {
	Title   string `json:"title" binding:"omitempty,min=1,max=200"`
	Content string `json:"content" binding:"omitempty,min=1"`
}

// CreatePost 创建文章
func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID.(uint),
	}

	if err := config.DB.Create(&post).Error; err != nil {
		utils.Logger.Error("创建文章失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建文章失败")
		return
	}

	// 加载关联的用户信息
	config.DB.Preload("User").First(&post, post.ID)

	utils.SuccessResponse(c, http.StatusCreated, "创建成功", post)
}

// GetPosts 获取文章列表
func GetPosts(c *gin.Context) {
	var posts []models.Post

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 查询文章列表，预加载用户信息
	if err := config.DB.Preload("User").
		Order("created_at desc").
		Limit(pageSize).
		Offset(offset).
		Find(&posts).Error; err != nil {
		utils.Logger.Error("获取文章列表失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取文章列表失败")
		return
	}

	// 获取总数
	var total int64
	config.DB.Model(&models.Post{}).Count(&total)

	utils.SuccessResponse(c, http.StatusOK, "获取成功", gin.H{
		"posts":     posts,
		"page":      page,
		"page_size": pageSize,
		"total":     total,
	})
}

// GetPost 获取单个文章详情
func GetPost(c *gin.Context) {
	postID := c.Param("id")

	var post models.Post
	if err := config.DB.Preload("User").Preload("Comments.User").First(&post, postID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "获取成功", post)
}

// UpdatePost 更新文章
func UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	userID, _ := c.Get("user_id")

	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 检查是否是文章作者
	if post.UserID != userID.(uint) {
		utils.ErrorResponse(c, http.StatusForbidden, "无权限修改此文章")
		return
	}

	var req UpdatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}

	if err := config.DB.Model(&post).Updates(updates).Error; err != nil {
		utils.Logger.Error("更新文章失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "更新文章失败")
		return
	}

	// 重新加载文章数据
	config.DB.Preload("User").First(&post, postID)

	utils.SuccessResponse(c, http.StatusOK, "更新成功", post)
}

// DeletePost 删除文章
func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	userID, _ := c.Get("user_id")

	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	// 检查是否是文章作者
	if post.UserID != userID.(uint) {
		utils.ErrorResponse(c, http.StatusForbidden, "无权限删除此文章")
		return
	}

	// 软删除
	if err := config.DB.Delete(&post).Error; err != nil {
		utils.Logger.Error("删除文章失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "删除文章失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "删除成功", nil)
}
