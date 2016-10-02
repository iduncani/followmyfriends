package followmyfriends

type SegmentDataCollector struct {
	data Data
	segmentLoader SegmentLoader
}

type SegmentLoader interface {
	loadSegments(activityId int64) []int64
}

func (collector *SegmentDataCollector) getSegmentData(activityId int64) *Data {

	return &collector.data
}