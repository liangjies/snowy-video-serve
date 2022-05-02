package model

type VideosHistory struct {
	ID      uint   `json:"id" gorm:"column:id"`
	UserID  uint   `json:"userId" gorm:"column:user_id"`
	VideoID uint64 `json:"videoId" gorm:"column:video_id"`
	Nums    uint   `json:"nums" gorm:"column:nums"`
}
