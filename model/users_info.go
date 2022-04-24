package model

import (
	"snowy-video-serve/global"

	"github.com/satori/go.uuid"
)

type UsersInfo struct {
	global.SYS_MODEL
	UUID              uuid.UUID `json:"-" gorm:"column:uuid;comment:用户UUID"`           // 用户UUID
	Username          string    `json:"username" gorm:"column:username;comment:用户登录名"` // 用户登录名
	NickName          string    `json:"nickname" gorm:"column:nickname;comment:用户昵称"`  // 用户昵称
	Avatar            string    `json:"avatar" gorm:"column:avatar;comment:用户昵称"`      // 用户头像
	Password          string    `json:"-"  gorm:"column:password;comment:用户登录密码"`      // 用户登录密码
	FansCounts        uint      `json:"fansCounts"  gorm:"column:fans_counts;"`
	FollowCounts      uint      `json:"followCounts"  gorm:"column:follow_counts;"`
	ReceiveLikeCounts uint      `json:"receiveLikeCounts"  gorm:"column:receive_like_counts;"`
	Gender            uint      `json:"gender"  gorm:"column:gender;"` //性别
	BackgroundImage   string    `json:"backgroundImage"  gorm:"column:background_image;"`
	Signature         string    `json:"signature"  gorm:"column:signature;"`
}
