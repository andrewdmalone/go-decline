package forecast

type EndConditionType int

const (
	Rate EndConditionType = iota
	Time
	CumulativeVolume
	DeclineRate
)

type EndCondition struct {
	EndConditionType EndConditionType
	Value            float64
}

type ForecastSegment struct {
	Model        Model
	EndCondition EndCondition
}

func NewForecastSegment(Model Model, EndCondition EndCondition) *ForecastSegment {
	return &ForecastSegment{Model: Model, EndCondition: EndCondition}
}

func (fs *ForecastSegment) DaysToSegmentEnd() float64 {
	return 0.0
}
