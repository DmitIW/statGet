package randomDist

import (
	"math"
	"math/rand"
)

func NormalX(mean float64, std float64) float64 {
	x := rand.NormFloat64()
	return x*std + mean
}

func ABSNormal(mean float64, std float64) float64 {
	return math.Abs(NormalX(mean, std))
}
