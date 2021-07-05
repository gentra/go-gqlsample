package resolver

import (
	"errors"

	"github.com/gentra/go-gqlsample/internal/enum"
	"github.com/gentra/go-gqlsample/internal/fetcher"
	"github.com/graphql-go/graphql"
)

type Lift struct {
}

func NewLift() *Lift {
	return &Lift{}
}

func (l *Lift) AllLifts() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return fetcher.GetLifts(), nil
	}
}

func (l *Lift) SetLiftStatus() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		id, ok := p.Args["id"].(string)
		if !ok {
			return nil, errors.New("error casting input id")
		}

		status, ok := p.Args["status"].(enum.LiftStatus)
		if !ok {
			return nil, errors.New("error casting input status")
		}

		data, err := fetcher.SetLiftStatus(id, status)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
}
