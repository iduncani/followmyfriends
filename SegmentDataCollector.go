package followmyfriends

type SegmentDataCollector struct {
	data          Data
	segmentLoader SegmentLoader
}

type SegmentLoader interface {
	loadSegments(activityId int64) []int64
}

func (collector *SegmentDataCollector) getSegmentData(activityId int64) *Data {
	segmentIds := collector.segmentLoader.loadSegments(activityId)
	for _, id := range segmentIds {
		segmentData := SegmentData{id}
		collector.data.segments = append(collector.data.segments, &segmentData)
	}
	return &collector.data
}
