package model

type Videos struct {
	ID      uint `json:"id" gorm:"column:id"`
	UserID  uint `json:"user_id" gorm:"column:user_id"`
	VideoID uint `json:"video_id" gorm:"column:video_id"`
}
