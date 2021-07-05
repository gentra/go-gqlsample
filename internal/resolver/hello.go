package resolver

import "github.com/graphql-go/graphql"

type Hello struct {
}

func NewHello() *Hello {
	return &Hello{}
}

func (h *Hello) Hello() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		return "Hello, world!", nil
	}
}
