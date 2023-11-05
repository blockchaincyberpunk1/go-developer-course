package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}
}

func TestSubtract(t *testing.T) {
	result := Subtract(5, 2)
	expected := 3
	if result != expected {
		t.Errorf("Subtract(5, 2) returned %d, expected %d", result, expected)
	}
}

func TestAddTableDriven(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{10, -5, 5},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) returned %d, expected %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
