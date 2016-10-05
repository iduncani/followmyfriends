package followmyfriends

import (
	"github.com/strava/go.strava"
	"log"
)

type StravaActivityLoader struct {
	accessToken string
}

func (loader *StravaActivityLoader) lastActivityIdForAthlete(athleteId int64) int64 {
	client := strava.NewClient(loader.accessToken)
	activities, err := strava.NewAthletesService(client).
		ListActivities(athleteId).
		PerPage(1).Do()
	if err != nil {
		log.Printf("Unable to find acitivites for %d. Error was %s,", athleteId, err.Error())
		return -1
	}
	if len(activities) == 0 {
		return -1
	} else {
		return activities[0].Id
	}
}
