package resolver

import (
	"github.com/gentra/go-gqlsample/internal/data"
	"github.com/gentra/go-gqlsample/internal/enum"
	"github.com/graph-gophers/graphql-go"
)

type Trail struct {
	Data data.Trail
}

func NewTrails(data []data.Trail) []*Trail {
	trails := make([]*Trail, len(data))
	for index := range data {
		trails[index] = NewTrail(data[index])
	}
	return trails
}

func NewTrail(data data.Trail) *Trail {
	return &Trail{Data: data}
}

func (t *Trail) ID() graphql.ID {
	return graphql.ID(t.Data.ID)
}

func (t *Trail) Name() string {
	return t.Data.Name
}

func (t *Trail) Status() *enum.TrailStatus {
	return &t.Data.Status
}

func (t *Trail) Difficulty() string {
	return t.Data.Difficulty
}

func (t *Trail) Groomed() bool {
	return t.Data.Groomed
}

func (t *Trail) Trees() bool {
	return t.Data.Trees
}

func (t *Trail) Night() bool {
	return t.Data.Night
}
