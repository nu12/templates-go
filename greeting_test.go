package main

import "testing"

func TestHelloTo(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", "Hello, !"},
		{"single name", "World", "Hello, World!"},
		{"another name", "John", "Hello, John!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helloTo(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
