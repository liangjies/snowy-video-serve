package ffmpeg

import (
	"snowy-video-serve/global"
	"snowy-video-serve/model"

	"snowy-video-serve/core"
	"snowy-video-serve/initialize"
)

func init() {
	core.Viper("../config.yaml")      // 初始化Viper
	global.SYS_LOG = core.Zap()       // 初始化zap日志库
	global.SYS_DB = initialize.Gorm() // gorm连接数据库
}

func VideoData() (videoList []model.Videos, err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	db = db.Where("is_local=1 AND (cover_path IS NULL OR cover_path='')")
	err = db.Find(&videoList).Error
	return videoList, err
}

func UpdateCover(video model.Videos) (err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	db = db.Where("id=?", video.ID)
	err = db.Select("cover_path").Save(video).Error
	return
}

func VideoDataDuration() (videoList []model.Videos, err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	db = db.Where("is_local=1 AND video_seconds=0")
	err = db.Find(&videoList).Error
	return videoList, err
}

func UpdateDuration(video model.Videos) (err error) {
	db := global.SYS_DB.Model(&model.Videos{})
	db = db.Where("id=?", video.ID)
	err = db.Select("video_seconds").Save(video).Error
	return
}
