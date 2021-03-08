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
		"lowercase alphabet only source": {
			source: "hhfjjqeduenmckebdjsxkwdxsddwd",
			length: 10,
		},
		"numeric string only source": {
			source: "1234567890",
			length: 30,
		},
		"lowercase alpha-numeric source": {
			source: "1234567890qwertyuiopasdfghjklzxcvbnm",
			length: 40,
		},
		"lowercase alpha-numeric and characters source": {
			source: "1234567890qwertyuiopasdfghjklzxcvbnm!@#$%^&*()~`_+{}[]|\\;:'\",.<>/?",
			length: 100,
		},
		"all cases alpha-numeric and characters source": {
			source: "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM!@#$%^&*()~`_+{}[]|\\;:'\",.<>/?",
			length: 200,
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
