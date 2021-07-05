package main

import (
	"log"
	"net/http"

	"github.com/gentra/go-gqlsample/internal"
	"github.com/gentra/go-gqlsample/internal/resolver"
)

func main() {
	helloResolver := resolver.NewHello()
	liftResolver := resolver.NewLift()
	trailResolver := resolver.NewTrail()
	schemaWrapper := internal.NewSchemaWrapper(helloResolver, liftResolver, trailResolver)

	err := schemaWrapper.Init()
	if err != nil {
		log.Fatalf("unable to parse schema %+v", err)
	}

	http.Handle("/graphql", internal.NewHandler(schemaWrapper).Handle())
	http.Handle("/ui", http.HandlerFunc(internal.GraphiQLHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
