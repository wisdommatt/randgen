package randgen

import (
	"strconv"
	"strings"
	"testing"
)

func TestGenerateNumberFromSource(t *testing.T) {
	testCases := map[string]struct {
		source int
		length int
	}{
		"first test case": {
			source: 1234567890,
			length: 5,
		},
		"second test case": {
			source: 1234567890,
			length: 10,
		},
	}
	numberGenerator := NewNumberGenerator()
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			generatedNumber := numberGenerator.GenerateFromSource(testCase.source, testCase.length)
			if len(strconv.Itoa(generatedNumber)) != testCase.length {
				t.Fatalf("Generated number length should be %d, but got %d", testCase.length, len(strconv.Itoa(generatedNumber)))
			}
			sourceStr := strconv.Itoa(testCase.source)
			for _, elem := range strconv.Itoa(generatedNumber) {
				if !strings.Contains(sourceStr, string(elem)) {
					t.Fatalf("Generated number %s should not contain %s", strconv.Itoa(generatedNumber), string(elem))
				}
			}
		})
	}
}
