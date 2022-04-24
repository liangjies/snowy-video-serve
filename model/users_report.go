package model

import (
	"time"
)

type UsersReport struct {
	ID          uint      `json:"id" gorm:"column:id"`
	DealVideoID string    `json:"dealVideoId" gorm:"column:deal_video_id"`
	DealUserID  uint      `json:"dealUserId" gorm:"column:deal_user_id"`
	Title       string    `json:"title" gorm:"column:title"`
	Content     string    `json:"content"  gorm:"column:content"`
	UserID      uint      `json:"userId"  gorm:"column:user_id"`
	CreateDate  time.Time `json:"createDate"  gorm:"column:create_date;"`
}
