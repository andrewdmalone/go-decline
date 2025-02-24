package models

type Model interface {
	Rate(dt float64) float64
	CumulativeVolume(dt float64) float64
	DeclineRate(dt float64) float64
}
