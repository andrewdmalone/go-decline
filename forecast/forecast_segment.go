package forecast

type EndConditionType int

const (
	DeclineRate EndConditionType = iota
	Rate
	Time
	CumulativeVolume
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
