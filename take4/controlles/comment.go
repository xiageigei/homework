package controllers

import (
	"net/http"

	"homework/take4/config"
	"homework/take4/models"
	"homework/take4/utils"

	"github.com/gin-gonic/gin"
)

// CreateCommentRequest 创建评论请求结构
type CreateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// CreateComment 创建评论
func CreateComment(c *gin.Context) {
	var req CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "请求参数错误: "+err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	// 检查文章是否存在
	var post models.Post
	if err := config.DB.First(&post, req.PostID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	comment := models.Comment{
		Content: req.Content,
		UserID:  userID.(uint),
		PostID:  req.PostID,
	}

	if err := config.DB.Create(&comment).Error; err != nil {
		utils.Logger.Error("创建评论失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "创建评论失败")
		return
	}

	// 加载关联的用户信息
	config.DB.Preload("User").First(&comment, comment.ID)

	utils.SuccessResponse(c, http.StatusCreated, "评论成功", comment)
}

// GetCommentsByPost 获取文章的所有评论
func GetCommentsByPost(c *gin.Context) {
	postID := c.Param("post_id")

	// 检查文章是否存在
	var post models.Post
	if err := config.DB.First(&post, postID).Error; err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "文章不存在")
		return
	}

	var comments []models.Comment
	if err := config.DB.Preload("User").
		Where("post_id = ?", postID).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		utils.Logger.Error("获取评论列表失败: ", err)
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取评论列表失败")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "获取成功", comments)
}
