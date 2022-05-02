package model

import (
	"time"
)

type Videos struct {
	ID           uint64    `json:"id" gorm:"column:id"`
	UserID       uint      `json:"userId" gorm:"column:user_id"`
	AudioID      string    `json:"audioId" gorm:"column:audio_id"`
	VideoDesc    string    `json:"videoDesc" gorm:"column:video_desc"`
	VideoPath    string    `json:"videoPath" gorm:"column:video_path"`
	VideoSeconds float32   `json:"videoSeconds" gorm:"column:video_seconds"`
	VideoWidth   int       `json:"videoWidth" gorm:"column:video_width"`
	VideoHeight  int       `json:"VideoHeight" gorm:"column:video_height"`
	CoverPath    string    `json:"coverPath" gorm:"column:cover_path"`
	LikeCounts   uint64    `json:"likeCounts" gorm:"column:like_counts"`
	Status       int       `json:"status" gorm:"column:status"`
	CreateTime   time.Time `json:"createTime" gorm:"column:create_time"`
	IsLocal      bool      `json:"isLocal" gorm:"column:is_local"`
}
