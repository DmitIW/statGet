package randomDist

import "math"

func Poisson(k float64, lambda float64) float64 {
	eps := math.Exp(-1.0 * lambda)
	if k != 0 {
		numerator := math.Pow(lambda, k)
		denominator := float64(Factorial(uint64(k)))
		eps *= numerator / denominator
	}
	return eps
}
