package response

import (
	"snowy-video-serve/model"
)

type UsersInfoResponse struct {
	model.UsersInfo
	IsFollow bool `json:"follow" gorm:"column:is_follow`
}

type LoginResponse struct {
	User      model.UsersInfo `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}

type UserLikeResponse struct {
	Data bool `json:"data"`
}

type QueryFollowsResponse struct {
	ID        uint   `json:"id" gorm:"column:id`
	Avatar    string `json:"avatar" gorm:"column:avatar`
	Nickname  string `json:"nickname" gorm:"column:nickname`
	Signature string `json:"signature" gorm:"column:signature`
}

type QueryFansResponse struct {
	ID        uint   `json:"id" gorm:"column:id`
	Avatar    string `json:"avatar" gorm:"column:avatar`
	Nickname  string `json:"nickname" gorm:"column:nickname`
	Signature string `json:"signature" gorm:"column:signature`
}
