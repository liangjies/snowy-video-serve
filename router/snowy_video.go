package router

import (
	"snowy-video-serve/api/v1"

	"github.com/gin-gonic/gin"
)

func InitVideoRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("video")
	{
		UserRouter.POST("uploadAvatar", v1.UploadAvatar) //用户上传头像
	}
}
