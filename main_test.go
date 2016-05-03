package main

import "testing"

func TestDetectInputFormat(t *testing.T) {
	var input_test_cases = []struct {
		input    []byte
		expected string
	}{
		{[]byte("---\nhello world"), "yaml"},
	}

	for _, test_case := range input_test_cases {
		actual := detectInputFormat(test_case.input)
		if actual != test_case.expected {
			t.Errorf("Expected %q, Got %q (input: %q)", test_case.expected, actual, test_case.input)
		}
	}
}
