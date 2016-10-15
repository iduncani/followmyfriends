package followmyfriends

type Follower struct {
	input         FollowFriendInput
	printer       FollowFriendOutput
	segmentFinder SegmentFinder
}

type FollowFriendInput interface {
	athleteId() int64
}

type FollowFriendOutput interface {
}

type SegmentFinder interface {
	loadSegmentsForLastActivity(athleteId int64) *Data
}

type Data struct {
	segments []*SegmentData
}

type SegmentData struct {
	Id       int64
	runCount int32
}

func (follower *Follower) Follow() {
	follower.segmentFinder.loadSegmentsForLastActivity(follower.input.athleteId())
}
