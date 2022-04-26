package model

import (
	"time"
)

type Comments struct {
	ID              uint      `json:"id" gorm:"column:id"`
	ToUserID        uint      `json:"toUserId" gorm:"column:to_user_id"`
	FatherCommentID uint      `json:"fatherCommentId" gorm:"column:father_comment_id"`
	VideoID         string    `json:"videoId" gorm:"column:video_id"`
	FromUserID      uint      `json:"fromUserId" gorm:"column:from_user_id"`
	Comment         string    `json:"comment" gorm:"column:comment"`
	CreateDate      time.Time `json:"createDate" gorm:"column:create_date"`
}
