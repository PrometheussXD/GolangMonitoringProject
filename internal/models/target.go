package models

type Target struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval int    `json:"interval"`
}
