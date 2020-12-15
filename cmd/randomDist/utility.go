package randomDist

const (
	factLIM uint64 = 41
)

var (
	fact [factLIM]uint64
)

func Factorial(x uint64) (res uint64) {
	if x > factLIM {
		x = factLIM
	} else if x <= 0 {
		return 1
	}

	if fact[x] == 0 {
		fact[x] = Factorial(x-1) * x
	}

	return fact[x]
}
