package request

type UserLike struct {
	UserID         uint   `form:"userId" json:"userId"`
	VideoID        uint64 `form:"videoId" json:"videoId"`
	VideoCreaterID uint   `form:"videoCreaterId" json:"videoCreaterId"`
}
type SaveHistory struct {
	VideoID string `form:"videoId" json:"videoId"`
}

type QueryVideos struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize" form:"pageSize"` // 每页大小
	VideoID   string `json:"videoId"`
	VideoDesc string `json:"videoDesc" gorm:"column:video_desc"`
	UserID    uint   `json:"userId" gorm:"column:user_id"`
}
