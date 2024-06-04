package main

imort "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("expected 'Even', got '%s'", result)
	}
}
