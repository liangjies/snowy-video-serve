package v1

import (
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"snowy-video-serve/model/response"
	"snowy-video-serve/service"
	"snowy-video-serve/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags User
// @Summary 分页和搜索查询视频列表
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /show/showVideos [post]
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
