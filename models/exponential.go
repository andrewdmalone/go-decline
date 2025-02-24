package models

import "math"

type Exponential struct {
	Qi float64 // Initial rate
	Di float64 // Initial decline rate
}

// Creates an exponential model
func NewExponential(Qi, Di float64) *Exponential {
	return &Exponential{Qi: Qi, Di: Di}
}

// Calculates rate at a time dt days from the origin point
func (e Exponential) Rate(dt float64) float64 {
	return e.Qi * math.Exp(-e.Di*dt)
}
