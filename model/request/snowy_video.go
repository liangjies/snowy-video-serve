package request

type UserLike struct {
	UserID         uint   `form:"userId" json:"userId"`
	VideoID        uint64 `form:"videoId" json:"videoId"`
	VideoCreaterID uint   `form:"videoCreaterId" json:"videoCreaterId"`
}
type SaveHistory struct {
	VideoID string `form:"videoId" json:"videoId"`
}
