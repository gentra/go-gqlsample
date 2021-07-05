package internal

import (
	"github.com/gentra/go-gqlsample/internal/enum"
	"github.com/gentra/go-gqlsample/internal/resolver"
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	Schema        graphql.Schema
	helloResolver *resolver.Hello
	liftResolver  *resolver.Lift
	trailResolver *resolver.Trail
}

func NewSchemaWrapper(helloResolver *resolver.Hello, liftResolver *resolver.Lift,
	trailResolver *resolver.Trail) *SchemaWrapper {
	return &SchemaWrapper{helloResolver: helloResolver, liftResolver: liftResolver, trailResolver: trailResolver}
}

func (s *SchemaWrapper) Init() error {
	var LiftStatusEnum = graphql.NewEnum(graphql.EnumConfig{
		Name: "LiftStatus",
		Values: graphql.EnumValueConfigMap{
			"OPEN": &graphql.EnumValueConfig{
				Value: enum.LiftOpen,
			},
			"CLOSED": &graphql.EnumValueConfig{
				Value: enum.LiftClosed,
			},
			"HOLD": &graphql.EnumValueConfig{
				Value: enum.LiftHold,
			},
		},
		Description: "An enum describing the options for LiftStatus: OPEN, CLOSED, HOLD",
	})

	var TrailStatusEnum = graphql.NewEnum(graphql.EnumConfig{
		Name: "TrailStatus",
		Values: graphql.EnumValueConfigMap{
			"OPEN": &graphql.EnumValueConfig{
				Value: enum.TrailOpen,
			},
			"CLOSED": &graphql.EnumValueConfig{
				Value: enum.TrailClosed,
			},
		},
		Description: "An enum describing the options for TrailStatus: OPEN, CLOSED",
	})

	var TrailType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Trail",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"status": &graphql.Field{
				Type: TrailStatusEnum,
			},
			"difficulty": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"groomed": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"trees": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"night": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
		},
	})

	var LiftType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Lift",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"status": &graphql.Field{
				Type: LiftStatusEnum,
			},
			"capacity": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"night": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
			},
			"elevationGain": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"trailAccess": &graphql.Field{
				Type:        graphql.NewList(TrailType),
				Resolve:     s.trailResolver.TrailsByLift(),
				Description: "A list of trails that this Lift serves",
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"hello": &graphql.Field{
					Type:        graphql.NewNonNull(graphql.String),
					Resolve:     s.helloResolver.Hello(),
					Description: "A simple greeting to you, earthlings!",
				},
				"allLifts": &graphql.Field{
					Type:    graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(LiftType))),
					Resolve: s.liftResolver.AllLifts(),
					Description: "A Lift is a chairlift, gondola, tram, funicular, pulley, rope tow, or other means of" +
						" ascending a mountain.",
				},
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"setLiftStatus": &graphql.Field{
					Type:    graphql.NewNonNull(LiftType),
					Resolve: s.liftResolver.SetLiftStatus(),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(graphql.ID),
						},
						"status": &graphql.ArgumentConfig{
							Type: graphql.NewNonNull(LiftStatusEnum),
						},
					},
				},
			},
		}),
		Types: []graphql.Type{
			LiftStatusEnum,
			TrailStatusEnum,
		},
	})
	if err != nil {
		return err
	}

	s.Schema = schema
	return nil
}
