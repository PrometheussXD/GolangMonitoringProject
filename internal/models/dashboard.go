package models

import "time"

type DashboardTarget struct {
	ID uint

	Name string

	URL string

	IsUp bool

	ResponseTime int64

	LastCheck time.Time

	Uptime float64
}
