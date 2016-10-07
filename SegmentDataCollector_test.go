package followmyfriends

import "testing"

func TestSingleSegmentRecorded(t *testing.T) {
	mockSegmentLoader := MockSegmentLoader{segmentsToReturn: []int64{2}}
	testObject := NewSegmentDataCollector(&mockSegmentLoader)

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

func TestTwoSegmentsRecorded(t *testing.T) {
	mockSegmentLoader := MockSegmentLoader{segmentsToReturn: []int64{3, 4}}
	testObject := NewSegmentDataCollector(&mockSegmentLoader)

	result := testObject.getSegmentData(1)

	assertCorrectIdsFound(result, mockSegmentLoader, t)
}

func TestSegmentRunTwiceIsRecordedAsOneSegment(t *testing.T) {
	mockSegmentLoader := MockSegmentLoader{segmentsToReturn: []int64{3, 3}}
	testObject := NewSegmentDataCollector(&mockSegmentLoader)

	result := testObject.getSegmentData(1)

	if len(result.segments) != 1 {
		t.Errorf("Expected a single segment to be returned but got %d", len(result.segments))
	} else {
		segmentData := result.segments[0]
		if segmentData.Id != 3 {
			t.Errorf("Expected the segment ID of 3 but go %s", segmentData.Id)
		}
	}
}

func assertCorrectIdsFound(result *Data, loader MockSegmentLoader, t *testing.T) {
	if loader.capturedActivityId != 1 {
		t.Error("Expected the supplied activity ID to be given to segment loader but got ", loader.capturedActivityId)
	}
	if len(result.segments) != len(loader.segmentsToReturn) {
		t.Errorf("Expected the same number of segments (%d) as in the mocked loader but got %d", len(loader.segmentsToReturn), len(result.segments))
	} else {
		for _, resultSegmentData := range result.segments {
			if !isExpectedSegment(resultSegmentData, loader) {
				t.Errorf("A segment ID %s was returned that wasn't expected", resultSegmentData.Id)
			}

		}
	}
}

func isExpectedSegment(segmentData *SegmentData, loader MockSegmentLoader) bool {
	segmentId := segmentData.Id
	for _, expectedId := range loader.segmentsToReturn {
		if segmentId == expectedId {
			return true
		}
	}
	return false
}

type MockSegmentLoader struct {
	capturedActivityId int64
	segmentsToReturn   []int64
}

func (loader *MockSegmentLoader) loadSegments(activityId int64) []int64 {
	loader.capturedActivityId = activityId
	return loader.segmentsToReturn
}
