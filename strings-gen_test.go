package randgen

import (
	"strings"
	"testing"
)

func TestGenerateStringFromSource(t *testing.T) {
	testCases := map[string]struct {
		source string
		length int
	}{
		"lowercase alphabet source": {
			source: StringLowercaseAlphabetsSource,
			length: 10,
		},
		"uppercase alphabet source": {
			source: StringUppercaseAlphabetsSource,
			length: 20,
		},
		"all cases alphabet source": {
			source: StringAlphabetsSource,
			length: 30,
		},
		"symbols source": {
			source: StringSymbolsSource,
			length: 40,
		},
		"numeric string source": {
			source: StringNumbericSoure,
			length: 15,
		},
		"lowercase alpha-numeric source": {
			source: StringLowercaseAlphaNumericSource,
			length: 45,
		},
		"uppercase alpha-numeric source": {
			source: StringUppercaseAlphaNumericSource,
			length: 50,
		},
		"alphabet-numeric source": {
			source: StringAlphaNumericSource,
			length: 55,
		},
		"alpha-numeric symbols source": {
			source: StringAlphaNumericSymbolsSource,
			length: 60,
		},
	}
	stringGenerator := NewStringGenerator()
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			generatedString := stringGenerator.GenerateFromSource(testCase.source, testCase.length)
			if len(generatedString) != testCase.length {
				t.Fatalf("Expected generated string length to be %d, got %d", len(generatedString), testCase.length)
			}
			for _, elem := range generatedString {
				if !strings.Contains(testCase.source, string(elem)) {
					t.Fatalf("Generated string from %s should not contain %s", testCase.source, string(elem))
				}
			}
		})
	}
}
