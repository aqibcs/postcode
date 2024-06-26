package postcode

import (
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    error
	}{
		{
			description: "Canada, Port Colborne",
			input:       " L3K 4V9 ",
			expected:    nil,
		},
		{
			description: "India , New Delhi",
			input:       " 110002 ",
			expected:    nil,
		},
		{
			description: "Mexico,  Mexico City",
			input:       " 01030 ",
			expected:    nil,
		},
		{
			description: "Japan, Tokyo",
			input:       " 100-0012 ",
			expected:    nil,
		},
		{
			description: "W N St, Aberdeen AB24, UK",
			input:       " AB24 ",
			expected:    nil,
		},
		{
			description: "Empty postal code",
			input:       "",
			expected:    ErrEmpty,
		},
		{
			description: "Short postal code",
			input:       "A",
			expected:    ErrShort,
		},
		{
			description: "Inexistent country code",
			input:       "TY 1234",
			expected:    ErrInvalidCountry,
		},
		{
			description: "Inexistent postal code format",
			input:       "11111111111",
			expected:    ErrInvalidFormat,
		},
	}

	for _, testCase := range testCases {
		t.Logf("Validating `%s` (%s)", testCase.input, testCase.description)
		if err := Validate(testCase.input); err != testCase.expected {
			t.Errorf("expected: %v; got: %v", testCase.expected, err)
		}
	}
}
