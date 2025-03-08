package forecast

type Model interface {
	SetParams(params []float64) error
	Rate(dt float64) float64
	CumulativeVolume(dt float64) float64
	VolumeBetween(dtStart, dtEnd float64) float64
	DeclineRate(dt float64) float64
}
