package followmyfriends

import "testing"

func TestSingleSegmentRecorded(t *testing.T) {
	mockSegmentLoader := MockSegmentLoader{}
	testObject := SegmentDataCollector{segmentLoader: &mockSegmentLoader}

	result := testObject.getSegmentData(1)

	if mockSegmentLoader.capturedActivityId != 1 {
		t.Error("Expected the supplied activity ID to be given to segment loader but got ", mockSegmentLoader.capturedActivityId)
	}
	if len(result.segments) != 1 {
		t.Error("Expected a single mocked segment but got ", len(result.segments))
	} else {
		segmentData := result.segments[0]
		if segmentData.Id != 2 {
			t.Error("Expected the mock segment ID to be returned but got ", segmentData.Id)
		}
	}
}

type MockSegmentLoader struct {
	capturedActivityId int64
}

func (loader *MockSegmentLoader) loadSegments(activityId int64) []int64 {
	loader.capturedActivityId = activityId
	return []int64 {2}
}