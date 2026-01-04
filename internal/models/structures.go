package models

type Proc struct {
	PID        int     `json:"pid"`
	CPUPercent float64 `json:"cpu"`
	MemPercent float64 `json:"mem"`
}
