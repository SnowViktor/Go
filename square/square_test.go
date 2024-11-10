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

func TestCustomSquare(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		got := customSquare(100_000)
		want := 10_000_000_000

		check(t, got, want)
	})

	t.Run("negative number", func(t *testing.T) {
		got := customSquare(-100_000)
		want := 10_000_000_000

		check(t, got, want)
	})

	t.Run("0", func(t *testing.T) {
		got := customSquare(0)
		want := 0

		check(t, got, want)
	})

	t.Run("1", func(t *testing.T) {
		got := customSquare(1)
		want := 1

		check(t, got, want)
	})
}

func TestDirectSquare(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		got := directSquare(200_000)
		want := 40_000_000_000

		check(t, got, want)
	})

	t.Run("negative number", func(t *testing.T) {
		got := directSquare(-200_000)
		want := 40_000_000_000

		check(t, got, want)
	})

	t.Run("0", func(t *testing.T) {
		got := directSquare(0)
		want := 0

		check(t, got, want)
	})
}

func TestMathPow(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		got := mathPow(300_000)
		want := 90_000_000_000

		check(t, got, want)
	})

	t.Run("negative number", func(t *testing.T) {
		got := mathPow(-300_000)
		want := 90_000_000_000

		check(t, got, want)
	})

	t.Run("0", func(t *testing.T) {
		got := mathPow(0)
		want := 0

		check(t, got, want)
	})
}

func TestMeasureTime(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		gotResult, _ := measureTime(40, customSquare)
		wantResult := 1600

		check(t, gotResult, wantResult)
	})

	t.Run("negative number", func(t *testing.T) {
		gotResult, _ := measureTime(-40, directSquare)
		wantResult := 1600

		check(t, gotResult, wantResult)
	})

	t.Run("0", func(t *testing.T) {
		gotResult, _ := measureTime(0, mathPow)
		wantResult := 0

		check(t, gotResult, wantResult)
	})
}

func TestOutput(t *testing.T) {
	t.Run("positive number", func(t *testing.T) {
		Output(2024)
	})

	t.Run("negative number", func(t *testing.T) {
		Output(-2024)
	})

	t.Run("0", func(t *testing.T) {
		Output(0)
	})
}
