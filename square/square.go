package square

import (
	"fmt"
	"math"
	"time"
)

func customSquare(n int) int {
	if n < 0 {
		n = -n
	}

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	a := 1
	b := 3

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
	return int(math.Pow(float64(n), 2))
}

func measureTime(n int, squareFunc func(int) int) (int, int64) {
	start := time.Now()
	var result int
	if n < 0 {
		for i := -1; i >= n; i-- {
			result = squareFunc(i)
		}
	} else if n > 0 {
		for i := 1; i <= n; i++ {
			result = squareFunc(i)
		}
	} else {
		result = 0
	}
	duration := time.Since(start).Microseconds()
	return result, duration
}

func Output(n int) {
	customSquareResult, customSquareDuration := measureTime(n, customSquare)
	directSquareResult, directSquareDuration := measureTime(n, directSquare)
	mathPowResult, mathPowDuration := measureTime(n, mathPow)

	fmt.Printf("customSquare(%d):\n\tresult: %d\n\tduration: %d\n\n",
		n, customSquareResult, customSquareDuration)
	fmt.Printf("directSquare(%d):\n\tresult: %d\n\tduration: %d\n\n",
		n, directSquareResult, directSquareDuration)
	fmt.Printf("mathPow(%d):\n\tresult: %d\n\tduration: %d\n\n",
		n, mathPowResult, mathPowDuration)
}
