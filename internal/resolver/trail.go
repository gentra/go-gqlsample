package resolver

import (
	"github.com/gentra/go-gqlsample/internal/data"
	"github.com/gentra/go-gqlsample/internal/fetcher"
	"github.com/graphql-go/graphql"
)

type Trail struct {
}

func NewTrail() *Trail {
	return &Trail{}
}

func (t *Trail) TrailsByLift() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		lift, ok := p.Source.(data.Lift)
		if !ok {
			return nil, nil
		}
		return fetcher.GetTrailsByLiftID(lift.ID), nil
	}
}
