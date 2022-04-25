package request

type UserLike struct {
	UserID         uint   `form:"userId" json:"userId"`
	VideoID        string `form:"videoId" json:"videoId"`
	VideoCreaterID uint   `form:"videoCreaterId" json:"videoCreaterId"`
}
