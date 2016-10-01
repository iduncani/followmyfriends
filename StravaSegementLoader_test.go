package followmyfriends

import (
	"testing"
	"reflect"
)

func TestLoadSegments(t *testing.T) {
	if testing.Short() {
		t.Skip("Test makes external calls to strava so skip in short mode")
	}
	testObject := StravaSegementLoader{testToken}

	segmentIds := testObject.loadSegments(729828914)

	if !reflect.DeepEqual(segmentIds, []int64 {17900330759, 17900330735, 17900330771, 17900330787, 17900330937, 17900330801, 17900330852, 17900330861, 17900330925, 17900330841, 17900330869, 17900330883, 17900330914, 17900330895, 17900330952, 17900330960, 17900331041, 17900331069, 17900331082, 17900331056}) {
		t.Error("Expected known segment IDs but got ", segmentIds)
	}
}

func TestEmptySegmentsForInvalidActity(t *testing.T) {
	if testing.Short() {
		t.Skip("Test makes external calls to strava so skip in short mode")
	}
	testObject := StravaSegementLoader{testToken}

	segmentIds := testObject.loadSegments(9999999999)

	if len(segmentIds) != 0 {
		t.Error("Expected no segments but got ", segmentIds)
	}
}