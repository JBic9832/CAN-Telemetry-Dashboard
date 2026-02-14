package source

import (
	"math/rand/v2"

	"github.com/jbic9832/can_dashboard/telemetry"
)

type SimSource struct {
	freq int
}

func NewSimSourceSource(f int) *SimSource {
	return &SimSource{
		freq: f,
	}
}

// super basic for now
func (s *SimSource) GenerateFrame() *telemetry.Frame {
	rpm := 900.0
	speed := 0.0
	coolant := 20.0
	oil := 20.0

	rpm += (rand.Float64() - 0.5) * 400
	if rpm < 800 {
		rpm = 800
	}
	if rpm > 6500 {
		rpm = 6500
	}

	speed += (rpm/6500*120 - speed) * 0.05
	if speed < 0 {
		speed = 0
	}

	coolant += (90 - coolant) * 0.01
	if coolant < 20 {
		coolant = 20
	}

	oil += (95 - oil) * 0.008
	if oil < 20 {
		oil = 20
	}

	return &telemetry.Frame{
		RPM:         rpm,
		Speed:       speed,
		CoolantTemp: coolant,
		OilTemp:     oil,
	}
}
