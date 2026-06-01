package models

import "time"

type CheckResult struct {
	ID uint `gorm:"primaryKey" json:"id"`

	TargetID uint `json:"target_id"`

	StatusCode int `json:"status_code"`

	ResponseTime int64 `json:"response_time"`

	IsUp bool `json:"is_up"`

	CheckedAt time.Time `json:"checked_at"`
}
