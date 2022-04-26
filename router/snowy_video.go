package router

import (
	"snowy-video-serve/api/v1"

	"github.com/gin-gonic/gin"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("video")
	{
		UserRouter.POST("showVideos", v1.ShowVideos)                 //分页和搜索查询视频列表
		UserRouter.POST("showAllVideos", v1.ShowAllVideos)           //获取所有视频
		UserRouter.POST("userLike", v1.UserLike)                     //用户点赞
		UserRouter.POST("userUnlike", v1.UserUnlike)                 //用户取消点赞
		UserRouter.POST("showUserLike", v1.ShowUserLike)             //显示点赞过的视频
		UserRouter.POST("showMyFollowVideos", v1.ShowMyFollowVideos) //显示我关注的人发的视频
		UserRouter.POST("saveComment", v1.SaveComment)               //用户留言
		UserRouter.POST("getVideoComments", v1.GetVideoComments)     //获取视频用户留言
		UserRouter.POST("getAllComments", v1.GetAllComments)         //获取我发布的视频内其他用户的留言

	}
}
