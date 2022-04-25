package v1

import (
	"fmt"
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"snowy-video-serve/model/request"
	"snowy-video-serve/model/response"
	"snowy-video-serve/service"
	"snowy-video-serve/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Video
// @Summary 分页和搜索查询视频列表
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/showVideos [post]
func ShowVideos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page == 0 {
		page = 1
	}
	// 获取JSON数据
	var videos model.Videos
	_ = c.ShouldBindJSON(&videos)

	const PAGE_SIZE int = 6 // 每页分页的记录数
	if err, list, total := service.QueryAllVideos(utils.GetUserID(c), videos, page, PAGE_SIZE); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: PAGE_SIZE,
		}, "获取成功", c)
	}
}

// @Tags Video
// @Summary 获取所有视频
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/showVideos [post]
func ShowAllVideos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page == 0 {
		page = 1
	}
	var videos model.Videos

	const PAGE_SIZE int = 6 // 每页分页的记录数
	if err, list, total := service.QueryAllVideos(0, videos, page, PAGE_SIZE); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: PAGE_SIZE,
		}, "获取成功", c)
	}
}

// TODO 获取热搜词

// @Tags Video
// @Summary 用户点赞
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /video/userLike [post]
func UserLike(c *gin.Context) {
	// 获取GET数据
	var userLike request.UserLike
	_ = c.ShouldBindQuery(&userLike)
	fmt.Println(userLike)
	// 校验参数
	if userLike.VideoID == "" || userLike.UserID == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 用户点赞
	if err := service.UserLike(utils.GetUserID(c), userLike); err != nil {
		global.SYS_LOG.Error("点赞失败!", zap.Any("err", err))
		response.FailWithMessage("点赞失败", c)
	} else {
		response.Ok(c)
	}
}

// @Tags Video
// @Summary 用户取消点赞
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /video/userUnlike [post]
func UserUnlike(c *gin.Context) {
	// 获取GET数据
	var userLike request.UserLike
	_ = c.ShouldBindQuery(&userLike)
	// 校验参数
	if userLike.VideoID == "" || userLike.UserID == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 用户取消点赞
	if err := service.UserUnlike(utils.GetUserID(c), userLike); err != nil {
		global.SYS_LOG.Error("取消点赞失败!", zap.Any("err", err))
		response.FailWithMessage("取消点赞失败", c)
	} else {
		response.Ok(c)
	}
}

// @Tags Video
// @Summary 显示点赞过的视频
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/showUserLike [post]
func ShowUserLike(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page == 0 {
		page = 1
	}

	const PAGE_SIZE int = 6 // 每页分页的记录数
	if err, list, total := service.ShowUserLike(utils.GetUserID(c), page, PAGE_SIZE); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: PAGE_SIZE,
		}, "获取成功", c)
	}
}
