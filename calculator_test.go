package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 4, b: 2, want: 2},
		{a: 9, b: 5, want: 4},
		{a: 7, b: 9, want: -2},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 3, want: 6},
		{a: 3, b: 3, want: 9},
		{a: 4, b: -5, want: -20},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 10, b: 2, want: 5},
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)

		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}

		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()

	_, err := calculator.Divide(1, 0)

	if err == nil {
		t.Errorf("want error for invalid input, got nil")
	}
}
