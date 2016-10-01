package followmyfriends

type ActivitySegmentFinder struct {
	activityLoader ActivityLoader
	segmentLoader SegmentLoader
}

type ActivityLoader interface {
	lastActivityIdForAthlete(athleteId int64) int64
}

type SegmentLoader interface {
	loadSegments(activityId int64) *Data
}

func (finder *ActivitySegmentFinder) loadSegmentsForAthlete(athleteId int64) *Data {
	activityId := finder.activityLoader.lastActivityIdForAthlete(athleteId);
	segmentData := finder.segmentLoader.loadSegments(activityId);
	return segmentData
}