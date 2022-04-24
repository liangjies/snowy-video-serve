package service

import (
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"snowy-video-serve/model/response"
)

//@function: ShowVideos
//@description: 分页和搜索查询视频列表
//@param: id uint
//@return: err error, list interface{}, total int64
func QueryAllVideos(id uint, videos model.Videos, page int, PAGE_SIZE int) (err error, list interface{}, total int64) {
	limit := PAGE_SIZE
	offset := PAGE_SIZE * (page - 1)
	db := global.SYS_DB.Model(&model.Videos{})
	var videoList []response.ShowVideoResponse

	db = db.Select("videos.*, users_info.avatar, users_info.nickname").Joins("left join users_info on user_id = users_info.id").Where("status = 1")
	// 搜索
	if videos.VideoDesc != "" {
		db = db.Where("method LIKE ?", "%"+videos.VideoDesc+"%")
		// TODO 保存热搜词
	}
	// 搜索
	if videos.UserID != 0 {
		db = db.Where("user_id = ?", videos.UserID)
	}

	err = db.Count(&total).Error
	if err != nil {
		return err, videoList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("create_time desc").Find(&videoList).Error
	}

	return err, videoList, total
}
