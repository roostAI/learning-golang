package main

import (
	"testing"
)

func TestAdd_28d8c9c450(t *testing.T) {
	// Test case 1: Positive numbers
	a, b := 5, 10
	expected := 15
	result := Add(a, b)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
	}

	// Test case 2: Negative numbers
	a, b = -5, -10
	expected = -15
	result = Add(a, b)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
	}

	// Test case 3: Zero
	a, b = 0, 0
	expected = 0
	result = Add(a, b)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
	}

	// Test case 4: Positive and negative numbers
	a, b = 5, -10
	expected = -5
	result = Add(a, b)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
	}

	// Test case 5: Large numbers
	a, b = 1000000, 2000000
	expected = 3000000
	result = Add(a, b)
	if result != expected {
		t.Errorf("Add(%d, %d) = %d; want %d", a, b, result, expected)
	}
}
