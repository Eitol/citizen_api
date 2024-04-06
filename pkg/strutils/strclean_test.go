package strutils

import (
	"testing"
)

func TestRemoveAccents(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic",
			input:    "Mëtàl Hëàd",
			expected: "Metal Head",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "string with no accents",
			input:    "Heavy Metal",
			expected: "Heavy Metal",
		},
		{
			name:     "string with special characters",
			input:    "!@#$%^&*()_+",
			expected: "",
		},
		{
			name:     "string with numbers",
			input:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "string with letters ñ and Ñ",
			input:    "El Niño",
			expected: "El Niño",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveAccents(test.input)

			if result != test.expected {
				t.Errorf("Expected '%s', but got '%s'", test.expected, result)
			}
		})
	}
}

func BenchmarkRemoveAccents(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RemoveAccents("Përformänce Tëst! 1, 2, 3...")
	}
}
