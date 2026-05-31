package models

type Target struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name" binding:"required"`
	URL      string `json:"url" binding:"required,url"`
	Interval int    `json:"interval" binding:"required,min=1"`
}
