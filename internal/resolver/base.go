package resolver

import (
	"github.com/gentra/go-gqlsample/internal/enum"
	"github.com/gentra/go-gqlsample/internal/fetcher"
)

type BaseResolver struct{}

func (_ *BaseResolver) Hello() string { return "Hello, world!" }

func (_ *BaseResolver) AllLifts(args *struct {
	Status *enum.LiftStatus
}) []*Lift {
	return NewLifts(fetcher.GetLifts())
}

func (_ *BaseResolver) SetLiftStatus(args *struct {
	ID     string
	Status enum.LiftStatus
}) (*Lift, error) {
	data, err := fetcher.SetLiftStatus(args.ID, args.Status)
	if err != nil {
		return nil, err
	}

	return NewLift(data), nil
}
