package models

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
