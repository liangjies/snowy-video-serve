package service

import (
	"errors"
	"mime/multipart"
	"path"
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"snowy-video-serve/model/response"
	"snowy-video-serve/utils/upload"
	"strings"
	"time"

	"gorm.io/gorm"
)

//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func Upload(file model.FileUpload) error {
	return global.SYS_DB.Create(&file).Error
}

//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func UploadUserImage(header *multipart.FileHeader, userId uint, imgType string) (err error, file model.FileUpload) {
	// 文件后缀校验
	fileExt := strings.ToLower(path.Ext(header.Filename))
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".gif" && fileExt != ".jpeg" {
		return errors.New("上传失败!只允许png,jpg,gif,jpeg文件"), file
	}
	// 图片大小校验
	if header.Size/1024 > 5120 {
		return errors.New("上传失败!只允许小于5MB文件"), file
	}

	// 上传图片
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		return uploadErr, file
	}

	// 更新用户信息
	if imgType == "Avatar" {
		if err = UpdateUserInfo(model.UsersInfo{SYS_MODEL: global.SYS_MODEL{ID: userId}, Avatar: filePath}); err != nil {
			return err, file
		}
	} else if imgType == "Background" {
		if err = UpdateUserInfo(model.UsersInfo{SYS_MODEL: global.SYS_MODEL{ID: userId}, BackgroundImage: filePath}); err != nil {
			return err, file
		}
	}

	// 返回图像信息
	s := strings.Split(header.Filename, ".")
	f := model.FileUpload{
		Url:  filePath,
		Name: header.Filename,
		UID:  userId,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	return Upload(f), f

}

//@function: QueryUser
//@description: 查询用户信息
//@param: id uint
//@return: err error, user model.UsersInfo
func QueryUser(id uint) (err error, user model.UsersInfo) {
	db := global.SYS_DB.Model(&model.UsersInfo{})
	err = db.Where("id = ?", id).Find(&user).Error
	return err, user
}

//@function: UpdateUserInfo
//@description: 更新用户数据
//@param: user model.UsersInfo
//@return: err error
func UpdateUserInfo(user model.UsersInfo) (err error) {
	db := global.SYS_DB.Model(&model.UsersInfo{})
	err = db.Where("id = ?", user.ID).Updates(user).Error
	return err
}

//@function: QueryUserLike
//@description: 查询用户点赞信息
//@param: id uint
//@return: err error, list interface{}, total int64
func QueryUserLike(id uint, videoId string) (err error, userLikeVideo bool) {
	var userLikeVideos model.UsersLikeVideos
	db := global.SYS_DB.Model(&model.UsersLikeVideos{})
	err = db.Where("user_id = ? AND video_id=?", id, videoId).First(&userLikeVideos).Error
	// 记录不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false
	}
	return err, true
}

//@function: Follow
//@description: 关注用户
//@param: id uint
//@return: err error
func Follow(id uint, usersFan model.UsersFans) (err error) {
	var usersFans model.UsersFans
	db := global.SYS_DB.Model(&model.UsersFans{})
	// 首先查询有没有关注过
	if err = db.Where("user_id = ? AND fan_id=?", usersFan.FanID, id).First(&usersFans).Error; err == nil {
		return nil
	}
	// 开始事务
	tx := db.Begin()
	// 关注用户，这里是被动关系
	if err = tx.Create(&model.UsersFans{UserID: usersFan.FanID, FanID: id}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 关注用户增加关注数
	if err = tx.Model(&model.UsersInfo{}).Where("id = ?", id).Update("follow_counts", gorm.Expr("follow_counts + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 被关注用户增加粉丝数
	if err = tx.Model(&model.UsersInfo{}).Where("id = ?", usersFan.FanID).Update("fans_counts", gorm.Expr("fans_counts + 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 事务提交
	return tx.Commit().Error
}

//@function: UnFollow
//@description: 取消关注用户
//@param: id uint
//@return: err error
func UnFollow(id uint, usersFan model.UsersFans) (err error) {
	var usersFans model.UsersFans
	db := global.SYS_DB.Model(&model.UsersFans{})
	// 开始事务
	tx := db.Begin()
	// 取消关注，这里是被动关系
	if err = db.Where("user_id = ? AND fan_id=?", usersFan.FanID, id).Delete(&usersFans).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 关注用户减少关注数
	if err = tx.Model(&model.UsersInfo{}).Where("id = ?", id).Update("follow_counts", gorm.Expr("follow_counts - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 被关注用户减少粉丝数
	if err = tx.Model(&model.UsersInfo{}).Where("id = ?", usersFan.FanID).Update("fans_counts", gorm.Expr("fans_counts - 1")).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 事务提交
	return tx.Commit().Error
}

//@function: QueryFollows
//@description: 获取关注用户信息
//@param: id uint
//@return: err error, list interface{}, total int64
func QueryFollows(id uint) (err error, list interface{}, total int64) {
	var queryFollows []response.QueryFollowsResponse
	db := global.SYS_DB.Model(&model.UsersFans{})
	db = db.Joins("left join users_info on user_id = users_info.id").Where("fan_id = ?", id)
	err = db.Count(&total).Error
	if err != nil {
		return err, queryFollows, total
	} else {
		err = db.Scan(&queryFollows).Error
	}

	return err, queryFollows, total
}

//@function: QueryFans
//@description: 获取粉丝信息
//@param: id uint
//@return: err error, list interface{}, total int64
func QueryFans(id uint) (err error, queryFans []response.QueryFansResponse) {
	db := global.SYS_DB.Model(&model.UsersFans{})
	err = db.Joins("left join users_info on fan_id = users_info.id").Where("user_id = ?", id).Scan(&queryFans).Error
	return err, queryFans
}

//@function: ReportUser
//@description: 举报用户
//@param: id uint
//@return: err error, list interface{}, total int64
func ReportUser(id uint, usersReport model.UsersReport) (err error) {
	usersReport.UserID = id
	usersReport.CreateDate = time.Now()
	db := global.SYS_DB.Model(&model.UsersReport{})

	err = db.Create(&usersReport).Error
	return err
}
