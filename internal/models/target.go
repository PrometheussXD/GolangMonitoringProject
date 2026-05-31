package models

type Target struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval int    `json:"interval"`
}
