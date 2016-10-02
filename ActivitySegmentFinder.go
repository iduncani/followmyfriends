package followmyfriends

type ActivitySegmentFinder struct {
	activityLoader ActivityLoader
	segmentLoader SegmentDataGenerator
}

type ActivityLoader interface {
	lastActivityIdForAthlete(athleteId int64) int64
}

type SegmentDataGenerator interface {
	getSegmentData(activityId int64) *Data
}

func (finder *ActivitySegmentFinder) loadSegmentsForAthlete(athleteId int64) *Data {
	activityId := finder.activityLoader.lastActivityIdForAthlete(athleteId);
	segmentData := finder.segmentLoader.getSegmentData(activityId);
	return segmentData
}