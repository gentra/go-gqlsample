package data

import "github.com/gentra/go-gqlsample/internal/enum"

type Trail struct {
	ID         string
	Name       string
	Status     enum.TrailStatus
	Difficulty string
	Groomed    bool
	Trees      bool
	Night      bool
}
