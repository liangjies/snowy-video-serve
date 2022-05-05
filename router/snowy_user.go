package router

import (
	"snowy-video-serve/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("uploadAvatar", v1.UploadAvatar)                   //用户上传头像
		UserRouter.POST("uploadBackgroundImage", v1.UploadBackgroundImage) //用户上传背景图片
		UserRouter.POST("updateNickName", v1.UpdateNickName)               //用户修改昵称
		UserRouter.POST("updateSignature", v1.UpdateSignature)             //用户修改个性签名
		UserRouter.POST("updateGender", v1.UpdateGender)                   //用户修改性别
		UserRouter.POST("queryUserLike", v1.QueryUserLike)                 //查询用户点赞信息
		UserRouter.POST("follow", v1.Follow)                               //关注用户
		UserRouter.POST("unFollow", v1.UnFollow)                           //取消关注用户
		UserRouter.POST("queryFollows", v1.QueryFollows)                   //获取关注用户信息
		UserRouter.POST("queryFans", v1.QueryFans)                         //获取粉丝信息
		UserRouter.POST("reportUser", v1.ReportUser)                       //举报用户
		UserRouter.POST("refreshToken", v1.RefreshToken)                   //更新token
		UserRouter.POST("query", v1.QueryUser)                             //查询用户信息
	}
}
