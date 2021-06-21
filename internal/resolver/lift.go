package resolver

import (
	"github.com/gentra/go-gqlsample/internal/data"
	"github.com/gentra/go-gqlsample/internal/enum"
	"github.com/gentra/go-gqlsample/internal/fetcher"
	"github.com/graph-gophers/graphql-go"
)

type Lift struct {
	Data data.Lift
}

func NewLifts(data []data.Lift) []*Lift {
	lifts := make([]*Lift, len(data))
	for index := range data {
		lifts[index] = NewLift(data[index])
	}
	return lifts
}

func NewLift(data data.Lift) *Lift {
	return &Lift{Data: data}
}

func (l *Lift) ID() graphql.ID {
	return graphql.ID(l.Data.ID)
}

func (l *Lift) Name() string {
	return l.Data.Name
}

func (l *Lift) Status() *enum.LiftStatus {
	return &l.Data.Status
}

func (l *Lift) Capacity() int32 {
	return l.Data.Capacity
}

func (l *Lift) Night() bool {
	return l.Data.Night
}

func (l *Lift) ElevationGain() int32 {
	return l.Data.ElevationGain
}

func (l *Lift) TrailAccess() []*Trail {
	return NewTrails(fetcher.GetTrailsByLiftID(l.Data.ID))
}
