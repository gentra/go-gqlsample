package fetcher

import (
	"fmt"

	"github.com/gentra/go-gqlsample/internal/data"
	"github.com/gentra/go-gqlsample/internal/enum"
)

var lifts []data.Lift

func GetLifts() []data.Lift {
	if len(lifts) == 0 {
		lifts = []data.Lift{
			{
				ID:            "1",
				Name:          "Lift-1",
				Status:        enum.LiftOpen,
				Capacity:      10,
				Night:         false,
				ElevationGain: 120,
			},
			{
				ID:            "2",
				Name:          "Lift-2",
				Status:        enum.LiftClosed,
				Capacity:      5,
				Night:         false,
				ElevationGain: 150,
			},
			{
				ID:            "3",
				Name:          "Lift-3",
				Status:        enum.LiftOpen,
				Capacity:      9,
				Night:         false,
				ElevationGain: 90,
			},
		}
	}
	return lifts
}

func SetLiftStatus(id string, status enum.LiftStatus) (data.Lift, error) {
	var updatedLift data.Lift
	for index := range GetLifts() {
		if lifts[index].ID == id {
			lifts[index].Status = status
			updatedLift = lifts[index]
		}
	}

	if updatedLift == (data.Lift{}) {
		return data.Lift{}, fmt.Errorf("lift with ID %s not found", id)
	}

	return updatedLift, nil
}
