package square

import (
	"testing"
)

func check(t *testing.T, got int, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func testSquareFunc(t *testing.T, squareFunc func(int) int, testCases map[string]struct {
	input int
	want  int
}) {
	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			got := squareFunc(tc.input)
			check(t, got, tc.want)
		})
	}
}

func TestSquareFunctions(t *testing.T) {
	testCases := map[string]struct {
		input int
		want  int
	}{
		"positive number": {100_000, 10_000_000_000},
		"negative number": {-100_000, 10_000_000_000},
		"0":               {0, 0},
		"1":               {1, 1},
	}

	squareFuncs := map[string]func(int) int{
		"customSquare":    customSquare,
		"directSquare":    directSquare,
		"mathPow":         mathPow,
		"iterativeSquare": iterativeSquare,
	}

	for _, squareFunc := range squareFuncs {
		testSquareFunc(t, squareFunc, testCases)
	}
}

func TestMeasureTime(t *testing.T) {
	testCases := map[string]struct {
		input int
		want  int
	}{
		"positive number": {40, 1600},
		"negative number": {-40, 1600},
		"0":               {0, 0},
		"1":               {1, 1},
	}

	for desc, tc := range testCases {
		t.Run(desc, func(t *testing.T) {
			gotResult, _ := measureTime(tc.input, customSquare)
			check(t, gotResult, tc.want)
		})
	}
}

func TestOutput(t *testing.T) {
	testCases := map[string]int{
		"positive number": 2024,
		"negative number": -2024,
		"0":               0,
	}

	for desc, input := range testCases {
		t.Run(desc, func(t *testing.T) {
			Output(input)
		})
	}
}
