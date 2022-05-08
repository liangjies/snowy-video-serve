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
	var queryVideos request.QueryVideos
	_ = c.ShouldBindJSON(&queryVideos)

	if queryVideos.Page == 0 {
		queryVideos.Page = 1
	}
	if queryVideos.PageSize == 0 { // 每页分页的记录数
		queryVideos.PageSize = 3
	}

	if err, list, total := service.QueryAllVideos(utils.GetUserID(c), queryVideos); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     queryVideos.Page,
			PageSize: queryVideos.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Video
// @Summary 获取所有视频
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/showAllVideos [post]
func ShowAllVideos(c *gin.Context) {
	var queryVideos request.QueryVideos
	_ = c.ShouldBindJSON(&queryVideos)

	if queryVideos.Page == 0 {
		queryVideos.Page = 1
	}
	if queryVideos.PageSize == 0 { // 每页分页的记录数
		queryVideos.PageSize = 3
	}

	if err, list, total := service.QueryAllVideos(utils.GetUserID(c), queryVideos); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     queryVideos.Page,
			PageSize: queryVideos.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Video
// @Summary 获取热搜词
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/hot [post]
func Hot(c *gin.Context) {
	if err, list := service.Hot(); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(list, "获取成功", c)
	}
}

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
	if userLike.VideoID == 0 || userLike.UserID == 0 {
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
	if userLike.VideoID == 0 || userLike.UserID == 0 {
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

// @Tags Video
// @Summary 显示我关注的人发的视频
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/showMyFollowVideos [post]
func ShowMyFollowVideos(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page == 0 {
		page = 1
	}

	const PAGE_SIZE int = 6 // 每页分页的记录数
	if err, list, total := service.ShowMyFollowVideos(utils.GetUserID(c), page, PAGE_SIZE); err != nil {
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
// @Summary 用户留言
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /video/saveComment [post]
func SaveComment(c *gin.Context) {
	// 获取GET数据
	var comments model.Comments
	_ = c.ShouldBindJSON(&comments)
	// 校验参数
	if comments.VideoID == 0 || comments.Comment == "" || utils.GetUserID(c) != comments.FromUserID {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 用户留言
	if err := service.SaveComment(utils.GetUserID(c), comments); err != nil {
		global.SYS_LOG.Error("评论失败!", zap.Any("err", err))
		response.FailWithMessage("评论失败", c)
	} else {
		response.Ok(c)
	}
}

// @Tags Video
// @Summary 获取视频用户留言
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/getVideoComments [post]
func GetVideoComments(c *gin.Context) {
	videoId, err := strconv.ParseUint(c.Query("videoId"), 10, 64)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page == 0 {
		page = 1
	}
	// 校验参数
	if videoId == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	const PAGE_SIZE int = 6 // 每页分页的记录数
	if err, list, total := service.GetVideoComments(utils.GetUserID(c), videoId, page, PAGE_SIZE); err != nil {
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
// @Summary 获取我发布的视频内其他用户的留言
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /video/GetAllComments [post]
func GetAllComments(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page == 0 {
		page = 1
	}

	const PAGE_SIZE int = 6 // 每页分页的记录数
	if err, list, total := service.GetAllComments(utils.GetUserID(c), page, PAGE_SIZE); err != nil {
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
// @Summary 保存播放记录
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /video/saveHistory [post]
func SaveHistory(c *gin.Context) {
	// 获取GET数据
	var saveHistory request.SaveHistory
	_ = c.ShouldBindJSON(&saveHistory)
	fmt.Println("videosHistory", saveHistory)
	videoID, err := strconv.ParseUint(saveHistory.VideoID, 10, 64)
	// 校验参数
	if err != nil || videoID == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 用户留言
	if err := service.SaveHistory(utils.GetUserID(c), videoID); err != nil {
		global.SYS_LOG.Error("保存失败!", zap.Any("err", err))
		response.FailWithMessage("保存失败", c)
	} else {
		response.Ok(c)
	}
}
