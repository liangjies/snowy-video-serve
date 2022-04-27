package model

type UsersFans struct {
	ID     uint `form:"id" json:"id" gorm:"column:id"`
	UserID uint `form:"userId" json:"userId" gorm:"column:user_id"`
	FanID  uint `form:"fanId" json:"fanId" gorm:"column:fan_id"`
}
