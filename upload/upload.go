package upload

import (
	"fmt"
	"log"
	"os"
	"snowy-video-serve/core"
	"snowy-video-serve/global"
	"snowy-video-serve/initialize"
	"snowy-video-serve/model"
	"snowy-video-serve/utils/upload"
)

// 初始化
func init() {
	core.Viper("../config.yaml")      // 初始化Viper
	global.SYS_LOG = core.Zap()       // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
}

// 获取要上传的视频
func VideoData() (videoList []model.Videos, err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	db = db.Where("is_local=1 AND cover_path IS NOT NULL AND cover_path IS NOT NULL")
	err = db.Find(&videoList).Error
	return videoList, err
}

func UpdateVideoData(video model.Videos) (err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	db = db.Where("id=?", video.ID)
	err = db.Updates(video).Error
	return err
}

// 上传视频到七牛
func UploadQiniu() {

	// 获取视频列表
	videoList, err := VideoData()
	if err != nil {
		fmt.Println("获取视频列表错误")
		return
	}

	var oss upload.OSS
	oss = &upload.Qiniu{}
	for _, video := range videoList {
		videoPath := fmt.Sprintf("../assets/video/%d.mp4", video.ID)
		coverPath := fmt.Sprintf("../assets/cover/%d.jpg", video.ID)
		videoName := fmt.Sprintf("%d.mp4", video.ID)
		coverName := fmt.Sprintf("%d.jpg", video.ID)

		// 视频
		videoOSSPath, _, uploadErr := oss.UploadLocalFile(videoPath, videoName, "video/")
		if uploadErr != nil {
			fmt.Println(uploadErr)
			continue
		}
		video.VideoPath = videoOSSPath

		// 封面
		coverOSSPath, _, uploadErr := oss.UploadLocalFile(coverPath, coverName, "cover/")
		if uploadErr != nil {
			fmt.Println(uploadErr)
			continue
		}
		video.CoverPath = coverOSSPath
		fmt.Println(video)
		// 写入数据库
		if err = UpdateVideoData(video); err != nil {
			fmt.Println(err)
			continue
		}

		// 删除本地文件
		// err = os.Remove(videoPath)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// err = os.Remove(coverPath)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		return
	}
}

func ReadDir() {
	dirname := "../assets/video/"

	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
