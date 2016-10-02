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
	loadSegmentsForAthlete(athleteId int64) *Data
}

type Data struct {
	segments []SegmentData
}

type SegmentData struct {
	Id int64
}

func (follower *Follower) Follow()  {
	follower.segmentFinder.loadSegmentsForAthlete(follower.input.athleteId())
}