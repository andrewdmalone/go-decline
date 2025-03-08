package forecast

import (
	"errors"
	"math"
)

type Hyperbolic struct {
	Qi float64 // Initial rate
	Di float64 // Initial decline rate
	B  float64 // b-factor
}

func NewHyperbolic(Qi, Di, B float64) *Hyperbolic {
	return &Hyperbolic{Qi: Qi, Di: Di, B: B}
}

func (h *Hyperbolic) SetParams(params []float64) error {
	if len(params) < 3 {
		return errors.New("hyperbolic decline requires 3 inputs - Qi (initial rate), Di (initial decline rate), and B (b-factor)")
	}
	if params[2] <= 0 {
		return errors.New("b-factor must be greater than zero")
	}
	h.Qi = params[0]
	h.Di = params[1]
	h.B = params[2]
	return nil
}

func (h *Hyperbolic) Rate(dt float64) float64 {
	return h.Qi * math.Pow(1+h.B*h.Di*dt, -1/h.B)
}

func (h *Hyperbolic) CumulativeVolume(dt float64) float64 {
	return h.VolumeBetween(0, dt)
}

func (h *Hyperbolic) VolumeBetween(dtStart, dtEnd float64) float64 {
	startingRate := h.Rate(dtStart)
	endingRate := h.Rate(dtEnd)
	firstTerm := math.Pow(startingRate, h.B) / ((1 - h.B) * h.DeclineRate(dtStart))
	secondTerm := (math.Pow(startingRate, 1-h.B) - math.Pow(endingRate, 1-h.B)) * daysPerYear
	return firstTerm * secondTerm
}

func (h *Hyperbolic) DeclineRate(dt float64) float64 {
	return h.Di / (1 + h.B*h.Di*dt)
}
