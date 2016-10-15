package followmyfriends

type SegmentDataCollector struct {
	segmentDataById map[int64]*SegmentData
	segmentLoader   SegmentLoader
	activityLoader  ActivityLoader
}

func NewSegmentDataCollector(segmentLoader SegmentLoader, activityLoader ActivityLoader) *SegmentDataCollector {
	return &SegmentDataCollector{make(map[int64]*SegmentData), segmentLoader, activityLoader}
}

type SegmentLoader interface {
	loadSegments(activityId int64) []int64
}

type ActivityLoader interface {
	lastActivityIdForAthlete(athleteId int64) int64
}

func (collector *SegmentDataCollector) loadSegmentsForLastActivity(athleteId int64) *Data {
	activityId := collector.activityLoader.lastActivityIdForAthlete(athleteId)
	return collector.getSegmentData(activityId)
}

func (collector *SegmentDataCollector) getSegmentData(activityId int64) *Data {
	segmentIds := collector.segmentLoader.loadSegments(activityId)
	for _, id := range segmentIds {
		collector.recordSegment(id)
	}
	allSegmentData := values(collector.segmentDataById)
	data := Data{allSegmentData}
	return &data
}

func (collector *SegmentDataCollector) recordSegment(segmentId int64) {
	segmentData, ok := collector.segmentDataById[segmentId]
	if !ok {
		segmentData = &SegmentData{segmentId, 1}
		collector.segmentDataById[segmentId] = segmentData
	} else {
		segmentData.runCount = segmentData.runCount + 1
	}
}

func values(fromMap map[int64]*SegmentData) []*SegmentData {
	values := make([]*SegmentData, len(fromMap))
	index := 0
	for _, segmentData := range fromMap {
		values[index] = segmentData
		index++
	}
	return values
}
