package model

type UsersLikeVideos struct {
	ID      uint   `json:"id" gorm:"column:id"`
	UserID  uint   `json:"user_id" gorm:"column:user_id"`
	VideoID string `json:"video_id" gorm:"column:video_id"`
}
