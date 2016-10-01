package followmyfriends

import "testing"

func TestFriendFollowerCanBeCreated(t *testing.T) {
	follower := Follower{input: &MockInput{}, printer: MockPrinter{}, segmentFinder: &MockSegmentFinder{}}
	if &follower == nil {
		t.Error("Expected to be able to create follower")
	}
}

func TestAtheleteIdPassedToDataLoader(t *testing.T) {
	dataLoader := MockSegmentFinder{}
	follower := Follower{input: &MockInput{}, printer: MockPrinter{}, segmentFinder: &dataLoader}

	follower.Follow()

	if dataLoader.athleteId != 1 {
		t.Error("Expected an athelete id of 1, got ", dataLoader.athleteId)
	}
}

type MockInput struct {

}

func (input *MockInput) athleteId() int64 {
	return 1;
}

type MockPrinter struct {

}

type MockSegmentFinder struct {
	athleteId int64
}

func (loader *MockSegmentFinder) loadSegmentsForAthlete(atheleteId int64) *Data {
	loader.athleteId = atheleteId
	return new(Data)
}