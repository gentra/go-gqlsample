package enum

type LiftStatus string

const (
	LiftOpen   LiftStatus = "OPEN"
	LiftClosed LiftStatus = "CLOSED"
	LiftHold   LiftStatus = "HOLD"
)
