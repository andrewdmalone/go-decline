package models

import (
	"errors"
	"math"
)

type Exponential struct {
	Qi float64 // Initial rate
	Di float64 // Initial decline rate
}

// Creates an exponential model
func NewExponential(Qi, Di float64) *Exponential {
	return &Exponential{Qi: Qi, Di: Di}
}

func (e *Exponential) SetParams(params []float64) error {
	if len(params) < 2 {
		return errors.New("Exponential decline requires 2 inputs - Qi (initial rate) and Di (initial decline rate)")
	}
	e.Qi = params[0]
	e.Di = params[1]
	// We place no restrictions on what Qi or Di can be as long as they are typed correctly. Negative values, zero, extremely high/low values are all valid.
	return nil
}

// Calculates rate at a time dt days from the origin point
func (e *Exponential) Rate(dt float64) float64 {
	return e.Qi * math.Exp(-e.Di*dt)
}

// Calculate the cumulative volume from the beginning of production to time dt
func (e *Exponential) CumulativeVolume(dt float64) float64 {
	return e.VolumeBetween(0, dt)
}

// Calculate the volume produced between dtStart and dtEnd
func (e *Exponential) VolumeBetween(dtStart, dtEnd float64) float64 {
	if e.Di == 0 {
		// Decline rate (Di) of zero implies constant rate
		return e.Qi * (dtStart - dtEnd)
	} else {
		// Nonzero decline rate (Di) handled normally
		return ((e.Rate(dtStart) - e.Rate(dtEnd)) / e.Di) * daysPerYear
	}
}

// Calculate the Decline Rate at a given point in time, dt
func (e *Exponential) DeclineRate(dt float64) float64 {
	return e.Di
}
