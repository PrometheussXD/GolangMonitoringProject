package models

import "time"

type DashboardTarget struct {
	ID           uint
	Name         string
	URL          string
	IsUp         bool
	ResponseTime int64
	LastCheck    time.Time
	Uptime       int
}

func (d DashboardTarget) FormattedLastCheck() string {
	if d.LastCheck.IsZero() {
		return "Never"
	}
	return d.LastCheck.Format("2006-01-02 15:04:05")
}
