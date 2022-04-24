package service

import (
	"errors"
	"snowy-video-serve/global"
	"snowy-video-serve/model"
	"snowy-video-serve/utils"

	"gorm.io/gorm"
)

//@function: Register
//@description: 用户注册
//@param: u model.UsersInfo
//@return: err error, userInter model.UsersInfo

func Register(u model.UsersInfo) (err error, userInter model.UsersInfo) {
	var user model.UsersInfo
	if !errors.Is(global.SYS_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.SYS_DB.Create(&u).Error
	return err, u
}

//@function: Login
//@description: 用户登录
//@param: u *model.UsersInfo
//@return: err error, userInter *model.UsersInfo

func Login(u *model.UsersInfo) (err error, userInter *model.UsersInfo) {
	var user model.UsersInfo
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.SYS_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

//@function: Login
//@description: 用户退出登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Logout(jwtList model.JwtBlacklist) (err error) {
	err = global.SYS_DB.Create(&jwtList).Error
	return
}

//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func ChangePassword(u *model.UsersInfo, newPassword string) (err error, userInter *model.UsersInfo) {
	var user model.UsersInfo
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.SYS_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.UsersInfo
//@return: err error, user model.UsersInfo

func SetUserInfo(reqUser model.UsersInfo) (err error, user model.UsersInfo) {
	err = global.SYS_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.UsersInfo

func FindUserById(id int) (err error, user *model.UsersInfo) {
	var u model.UsersInfo
	err = global.SYS_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.UsersInfo

func FindUserByUuid(uuid string) (err error, user *model.UsersInfo) {
	var u model.UsersInfo
	if err = global.SYS_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
