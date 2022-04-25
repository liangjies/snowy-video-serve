package service

import (
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"snowy-video-serve/model/request"
	"snowy-video-serve/model/response"

	"gorm.io/gorm"
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

//@function: UserLike
//@description: 用户点赞
//@param: id uint
//@return: err error, list interface{}, total int64
func UserLike(id uint, userLike request.UserLike) (err error) {
	db := global.SYS_DB.Model(&model.UsersLikeVideos{})
	// 开始事务
	tx := db.Begin()

	// 1.保存用户和视频的喜欢点赞关联关系表
	if err = tx.Create(&model.UsersLikeVideos{UserID: id, VideoID: userLike.VideoID}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 2.视频喜欢数量累加
	if err = tx.Model(&model.Videos{}).Where("id = ?", userLike.VideoID).Update("like_counts", gorm.Expr("like_counts + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 3.用户受喜欢数量累加
	var userID uint
	if err = tx.Model(&model.Videos{}).Select("user_id").Where("id = ?", userLike.VideoID).First(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Model(&model.UsersInfo{}).Where("id = ?", userID).Update("receive_like_counts", gorm.Expr("receive_like_counts + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 事务提交
	return tx.Commit().Error
}

//@function: UserUnlike
//@description: 用户取消点赞
//@param: id uint
//@return: err error, list interface{}, total int64
func UserUnlike(id uint, userLike request.UserLike) (err error) {
	var usersLikeVideos model.UsersLikeVideos
	db := global.SYS_DB.Model(&model.UsersLikeVideos{})
	// 开始事务
	tx := db.Begin()
	// 1.删除用户和视频的喜欢点赞关联关系表
	if err = tx.Where("user_id = ? AND video_id=?", id, userLike.VideoID).Delete(&usersLikeVideos).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 2.视频喜欢数量累减
	if err = tx.Model(&model.Videos{}).Where("id = ?", userLike.VideoID).Update("like_counts", gorm.Expr("like_counts - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 3.用户受喜欢数量累减
	var userID uint
	if err = tx.Model(&model.Videos{}).Select("user_id").Where("id = ?", userLike.VideoID).First(&userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Model(&model.UsersInfo{}).Where("id = ?", userID).Update("receive_like_counts", gorm.Expr("receive_like_counts - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 事务提交
	return tx.Commit().Error
}

//@function: ShowUserLike
//@description: 显示点赞过的视频
//@param: id uint
//@return: err error, list interface{}, total int64
func ShowUserLike(id uint, page int, PAGE_SIZE int) (err error, list interface{}, total int64) {
	limit := PAGE_SIZE
	offset := PAGE_SIZE * (page - 1)
	db := global.SYS_DB.Model(&model.UsersLikeVideos{})
	var videoList []response.ShowVideoResponse
	db = db.Where("users_like_videos.user_id=?", id)
	db = db.Select("videos.*, users_info.avatar, users_info.nickname").Joins("left join videos on videos.id = users_like_videos.video_id").Joins("left join users_info on users_like_videos.user_id = users_info.id").Where("status = 1")

	err = db.Count(&total).Error
	if err != nil {
		return err, videoList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Order("create_time desc").Find(&videoList).Error
	}

	return err, videoList, total
}
