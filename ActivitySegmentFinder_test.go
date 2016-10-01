package followmyfriends

import (
	"testing"
)

func TestSegmentsLoadedForLastActivity(t *testing.T) {
	var mockActivityLoader MockActivityLoader
	var mockSegmentLoader MockSegmentLoader
	activitySegmentFinder := ActivitySegmentFinder{&mockActivityLoader, &mockSegmentLoader}

	data := activitySegmentFinder.loadSegmentsForAthlete(5)

	if mockActivityLoader.capturedAthleteId != 5 {
		t.Error("Expected the athlete ID of 5 but got ", mockActivityLoader.capturedAthleteId)
	}
	if mockSegmentLoader.capturedActivityId != 1 {
		t.Error("Expected the activity ID of 1 but got ", mockSegmentLoader.capturedActivityId)
	}
	if data == nil {
		t.Error("No data was returned");
	}
}

type MockActivityLoader struct {
	capturedAthleteId int64
}

func (loader *MockActivityLoader) lastActivityIdForAthlete(athleteId int64) int64 {
	loader.capturedAthleteId = athleteId
	return 1;
}

type MockSegmentLoader struct {
	capturedActivityId int64
}

func (loader *MockSegmentLoader) loadSegments(activityId int64) *Data {
	loader.capturedActivityId = activityId
	return &Data{}
}