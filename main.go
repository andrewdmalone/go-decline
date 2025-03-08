package main

import (
	"fmt"

	"github.com/andrewdmalone/go-decline.git/forecast"
)

func main() {
	exp := forecast.NewExponential(1000, 0.1)
	for time := 0.0; time <= 5.0; time += 1.0 {
		rate := exp.Rate(time)
		fmt.Printf("Time: %.1f, Rate: %.2f\n", time, rate)
	}
}
