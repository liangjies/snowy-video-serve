package v1

import (
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

// @Tags User
// @Summary 用户上传头像
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadAvatar [post]
func UploadAvatar(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.SYS_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file := service.UploadUserImage(header, utils.GetUserID(c), "Avatar") // 文件上传后拿到文件路径
	if err != nil {
		global.SYS_LOG.Error("上传失败!", zap.Any("err", err))
		response.FailWithMessage("上传失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.FileUploadResponse{File: file}, "上传成功", c)
}

// @Tags User
// @Summary 用户上传背景图片
// @Produce  application/json
// @Param file File
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadAvatar [post]
func UploadBackgroundImage(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.SYS_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file := service.UploadUserImage(header, utils.GetUserID(c), "Background") // 文件上传后拿到文件路径
	if err != nil {
		global.SYS_LOG.Error("上传失败!", zap.Any("err", err))
		response.FailWithMessage("上传失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.FileUploadResponse{File: file}, "上传成功", c)
}

// @Tags User
// @Summary 用户修改昵称
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateNickName [POST]
func UpdateNickName(c *gin.Context) {
	// 获取原始数据
	row, err := c.GetRawData()
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 类型转换
	nickname := string(row)

	if nickname == "" {
		response.FailWithMessage("昵称不能为空", c)
		return
	}
	// 更新用户信息
	user := model.UsersInfo{SYS_MODEL: global.SYS_MODEL{ID: utils.GetUserID(c)}, NickName: nickname}
	if err := service.UpdateUserInfo(user); err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags User
// @Summary 用户修改个性签名
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateSignature [POST]
func UpdateSignature(c *gin.Context) {
	// 获取原始数据
	row, err := c.GetRawData()
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 类型转换
	signature := string(row)

	if signature == "" {
		response.FailWithMessage("签名不能为空", c)
		return
	}
	// 更新用户信息
	user := model.UsersInfo{SYS_MODEL: global.SYS_MODEL{ID: utils.GetUserID(c)}, Signature: signature}
	if err := service.UpdateUserInfo(user); err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags User
// @Summary 用户修改性别
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /user/updateGender [POST]
func UpdateGender(c *gin.Context) {
	// 获取原始数据
	row, err := c.GetRawData()
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 类型转换
	gender64, err := strconv.ParseUint(string(row), 10, 64)
	if err != nil {
		response.FailWithMessage("更新失败", c)
		return
	}
	gender := uint(gender64)
	// 判断是否空值
	if gender == 0 {
		response.FailWithMessage("更新失败", c)
		return
	}
	// 更新用户信息
	user := model.UsersInfo{SYS_MODEL: global.SYS_MODEL{ID: utils.GetUserID(c)}, Gender: gender}
	if err := service.UpdateUserInfo(user); err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags User
// @Summary 查询用户信息
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/query [post]
func QueryUser(c *gin.Context) {
	if err, user := service.QueryUser(utils.GetUserID(c)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.UsersInfoResponse{
			User: user,
		}, "获取成功", c)
	}
}

// @Tags User
// @Summary 查询用户点赞信息
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/queryUserLike [post]
func QueryUserLike(c *gin.Context) {
	// 获取JSON数据
	var userLike request.UserLike
	_ = c.ShouldBindJSON(&userLike)
	// 查询点赞信息
	if err, userLikeVideo := service.QueryUserLike(utils.GetUserID(c), userLike); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.UserLikeResponse{
			Data: userLikeVideo,
		}, "获取成功", c)
	}
}

// @Tags User
// @Summary 关注用户
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /user/follow [post]
func Follow(c *gin.Context) {
	// 获取JSON数据
	var usersFan model.UsersFans
	_ = c.ShouldBindJSON(&usersFan)
	// 关注用户
	if err := service.Follow(utils.GetUserID(c), usersFan); err != nil {
		global.SYS_LOG.Error("关注失败!", zap.Any("err", err))
		response.FailWithMessage("关注失败", c)
	} else {
		response.Ok(c)
	}
}

// @Tags User
// @Summary 取消关注用户
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /user/follow [post]
func UnFollow(c *gin.Context) {
	// 获取JSON数据
	var usersFan model.UsersFans
	_ = c.ShouldBindJSON(&usersFan)
	// 取消关注用户
	if err := service.UnFollow(utils.GetUserID(c), usersFan); err != nil {
		global.SYS_LOG.Error("取消关注失败!", zap.Any("err", err))
		response.FailWithMessage("取消关注失败", c)
	} else {
		response.Ok(c)
	}
}

// @Tags User
// @Summary 获取关注用户信息
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/queryFollows [post]
func QueryFollows(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("userId"), 10, 0)
	if err != nil || id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err, queryFollows := service.QueryFollows(uint(id)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(queryFollows, "获取成功", c)
	}
}

// @Tags User
// @Summary 获取粉丝信息
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/queryFans [post]
func QueryFans(c *gin.Context) {
	id, err := strconv.ParseUint(c.Query("userId"), 10, 0)
	if err != nil || id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err, queryFans := service.QueryFans(uint(id)); err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(queryFans, "获取成功", c)
	}
}

// @Tags User
// @Summary 举报用户
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"关注成功"}"
// @Router /user/reportUser [post]
func ReportUser(c *gin.Context) {
	// 获取JSON数据
	var usersReport model.UsersReport
	_ = c.ShouldBindJSON(&usersReport)

	if err := service.ReportUser(utils.GetUserID(c), usersReport); err != nil {
		global.SYS_LOG.Error("举报失败!", zap.Any("err", err))
		response.FailWithMessage("举报失败", c)
	} else {
		response.Ok(c)
	}
}
