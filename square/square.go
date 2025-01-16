package square

import (
	"fmt"
	"math"
	"time"
)

func customSquare(n int) int {
	if n < 0 {
		n = -n
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	a, b := 1, 3
	for i := 2; i <= n; i++ {
		a += b
		b += 2
	}
	return a
}

func directSquare(n int) int {
	return n * n
}

func mathPow(n int) int {
	return int(math.Pow(float64(n), 2.0))
}

func iterativeSquare(n int) int {
	if n < 0 {
		n = -n
	}
	result := 0
	for i := 0; i < n; i++ {
		result += n
	}
	return result
}

func measureTime(n int, squareFunc func(int) int) (int, int64) {
	start := time.Now()
	result := squareFunc(n)
	duration := time.Since(start).Microseconds()
	return result, duration
}

func Output(n int) {
	squareFuncs := map[string]func(int) int{
		"customSquare":    customSquare,
		"directSquare":    directSquare,
		"mathPow":         mathPow,
		"iterativeSquare": iterativeSquare,
	}

	for name, squareFunc := range squareFuncs {
		result, duration := measureTime(n, squareFunc)
		fmt.Printf("%s(%d):\n\tresult: %d\n\tduration: %d\n\n", name, n, result, duration)
	}
}
