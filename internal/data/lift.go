package data

import (
	"github.com/gentra/go-gqlsample/internal/enum"
)

type Lift struct {
	ID            string
	Name          string
	Status        enum.LiftStatus
	Capacity      int32
	Night         bool
	ElevationGain int32
}
