package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"Hello World", []string{"hello", "world"}},
		{"HELLO WORLD", []string{"hello", "world"}},
	}

	for _, c := range cases {
		var result []string = cleanInput(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("Expected length %d, got %d", len(c.expected), len(result))
			continue
		}
		for i := range result {
			if result[i] != c.expected[i] {
				t.Errorf("For input '%s', expected '%s' but got '%s'", c.input, c.expected[i], result[i])
			}
		}
	}
}