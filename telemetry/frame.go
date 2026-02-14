package telemetry

import "time"

type Frame struct {
	Timestamp   time.Time
	RPM         float64
	Speed       float64
	CoolantTemp float64
	OilTemp     float64
}

func CreateFrame(rpm float64, speed float64, coolantTemp float64, oilTemp float64) *Frame {
	return &Frame{
		Timestamp:   time.Now(),
		RPM:         rpm,
		Speed:       speed,
		CoolantTemp: coolantTemp,
		OilTemp:     oilTemp,
	}
}
