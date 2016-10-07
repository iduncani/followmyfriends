package followmyfriends

type SegmentDataCollector struct {
	segmentDataById map[int64]*SegmentData
	segmentLoader   SegmentLoader
}

func NewSegmentDataCollector(segmentLoader SegmentLoader) *SegmentDataCollector {
	return &SegmentDataCollector{make(map[int64]*SegmentData), segmentLoader}
}

type SegmentLoader interface {
	loadSegments(activityId int64) []int64
}

func (collector *SegmentDataCollector) getSegmentData(activityId int64) *Data {
	segmentIds := collector.segmentLoader.loadSegments(activityId)
	for _, id := range segmentIds {
		segmentData, ok := collector.segmentDataById[activityId]
		if !ok {
			segmentData = &SegmentData{id}
			collector.segmentDataById[id] = segmentData
		}
	}
	allSegmentData := make([]*SegmentData, len(collector.segmentDataById))
	index := 0
	for _, segmentData := range(collector.segmentDataById)  {
		allSegmentData[index] = segmentData
		index++
	}
	data := Data{ allSegmentData }
	return &data
}
