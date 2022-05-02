package response

import (
	"snowy-video-serve/model"
	"time"
)

type ShowVideoResponse struct {
	ID           string    `json:"id" gorm:"column:id"`
	UserID       uint      `json:"userId" gorm:"column:user_id"`
	AudioID      string    `json:"audioId" gorm:"column:audio_id"`
	VideoDesc    string    `json:"videoDesc" gorm:"column:video_desc"`
	VideoPath    string    `json:"videoPath" gorm:"column:video_path"`
	VideoSeconds float32   `json:"videoSeconds" gorm:"column:video_seconds"`
	VideoWidth   int       `json:"videoWidth" gorm:"column:video_width"`
	VideoHeight  int       `json:"VideoHeight" gorm:"column:video_height"`
	CoverPath    string    `json:"coverPath" gorm:"column:cover_path"`
	LikeCounts   uint64    `json:"likeCounts" gorm:"column:like_counts"`
	IsLocal      bool      `json:"isLocal" gorm:"column:is_local"`
	Status       int       `json:"status" gorm:"column:status"`
	CreateTime   time.Time `json:"createTime" gorm:"column:create_time"`
	Avatar       string    `json:"avatar" gorm:"column:avatar;comment:用户昵称"`     // 用户头像
	NickName     string    `json:"nickname" gorm:"column:nickname;comment:用户昵称"` // 用户昵称
	Isplay       bool      `json:"isplay" gorm:"column:isplay"`                  // APP
	PlayIng      bool      `json:"playIng" gorm:"column:playIng"`                // APP
	State        string    `json:"state" gorm:"column:state"`                    // APP
}

type VideoCommentResponse struct {
	model.Comments
	Avatar     string `json:"avatar" gorm:"column:avatar;comment:用户昵称"`     // 用户头像
	NickName   string `json:"nickname" gorm:"column:nickname;comment:用户昵称"` // 用户昵称
	ToNickName string `json:"toNickname" gorm:"column:to_nickname"`         // 用户昵称
}

type AllCommentResponse struct {
	model.Comments
	Avatar    string `json:"avatar" gorm:"column:avatar;comment:用户昵称"`     // 用户头像
	NickName  string `json:"nickname" gorm:"column:nickname;comment:用户昵称"` // 用户昵称
	UserID    uint   `json:"userId" gorm:"column:user_id"`
	VideoDesc string `json:"videoDesc" gorm:"column:video_desc"`
	CoverPath string `json:"coverPath" gorm:"column:cover_path"`
}
