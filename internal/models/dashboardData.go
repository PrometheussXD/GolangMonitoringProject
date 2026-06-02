package models

type DashboardData struct {
	Targets []DashboardTarget

	Count int

	UpCount int

	DownCount int

	Uptime float64
}
