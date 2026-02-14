package source

import "github.com/jbic9832/can_dashboard/telemetry"

type Source interface {
	GenerateFrame() *telemetry.Frame
}
