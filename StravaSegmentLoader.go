package followmyfriends

import (
	"github.com/strava/go.strava"
	"log"
)

type StravaSegementLoader struct {
	accessToken string
}

func (loader *StravaSegementLoader) loadSegments(activityId int64) []int64 {
	client := strava.NewClient(loader.accessToken)
	activityDetail, err := strava.NewActivitiesService(client).
		Get(activityId).
		IncludeAllEfforts().
		Do()
	if err != nil {
		log.Printf("Unable to find segments for %d. Error was %s,", activityId, err.Error())
		return []int64{}
	}
	segmentIds := make([]int64, len(activityDetail.SegmentEfforts))
	for index, segment := range activityDetail.SegmentEfforts {
		segmentIds[index] = segment.Id
	}
	return segmentIds
}
