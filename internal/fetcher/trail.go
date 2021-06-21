package fetcher

import (
	"github.com/gentra/go-gqlsample/internal/data"
	"github.com/gentra/go-gqlsample/internal/enum"
)

var trailsOfLift1 []data.Trail
var trailsOfLift2 []data.Trail

func GetTrailsByLiftID(liftID string) []data.Trail {
	switch liftID {
	case "1":
		return getTrailsOfLift1()
	case "2":
		return getTrailsOfLift2()
	default:
		return nil
	}
}

func getTrailsOfLift1() []data.Trail {
	if len(trailsOfLift1) == 0 {
		trailsOfLift1 = []data.Trail{
			{
				ID:         "1",
				Name:       "Trail-1",
				Status:     enum.TrailOpen,
				Difficulty: "Hard",
				Groomed:    false,
				Trees:      true,
				Night:      true,
			},
			{
				ID:         "2",
				Name:       "Trail-2",
				Status:     enum.TrailClosed,
				Difficulty: "Medium",
				Groomed:    false,
				Trees:      false,
				Night:      true,
			},
		}
	}
	return trailsOfLift1
}

func getTrailsOfLift2() []data.Trail {
	if len(trailsOfLift2) == 0 {
		trailsOfLift1 = []data.Trail{
			{
				ID:         "3",
				Name:       "Trail-3",
				Status:     enum.TrailClosed,
				Difficulty: "Hard",
				Groomed:    false,
				Trees:      true,
				Night:      false,
			},
			{
				ID:         "4",
				Name:       "Trail-4",
				Status:     enum.TrailOpen,
				Difficulty: "Medium",
				Groomed:    true,
				Trees:      false,
				Night:      true,
			},
			{
				ID:         "5",
				Name:       "Trail-5",
				Status:     enum.TrailOpen,
				Difficulty: "Medium",
				Groomed:    true,
				Trees:      true,
				Night:      true,
			},
		}
	}
	return trailsOfLift1
}
