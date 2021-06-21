package main

import (
	"github.com/gentra/go-gqlsample/internal/resolver"
	"log"
	"net/http"

	"github.com/gentra/go-gqlsample/internal"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schema := graphql.MustParseSchema(internal.Schema, &resolver.BaseResolver{})
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	http.Handle("/ui", http.HandlerFunc(internal.GraphiQLHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
