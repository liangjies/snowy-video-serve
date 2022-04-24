package model

type UsersFans struct {
	ID     uint `json:"id" gorm:"column:id"`
	UserID uint `json:"userId" gorm:"column:user_id"`
	FanID  uint `json:"fanId" gorm:"column:fan_id"`
}
