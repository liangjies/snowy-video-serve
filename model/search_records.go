package model

type SearchRecords struct {
	ID      uint   `json:"id" gorm:"column:id"`
	Content string `json:"content" gorm:"column:content"`
}
