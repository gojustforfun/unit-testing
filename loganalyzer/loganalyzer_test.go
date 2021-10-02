package loganalyzer_test

import (
	"testing"

	. "github.com/gojustforfun/unit-testing/loganalyzer"
	"github.com/stretchr/testify/assert"
)

func TestIsValidLogFileName(t *testing.T) {
	testcases := map[string]struct {
		fileName string
		expected bool
	}{
		"BadExtension_ShouldReturnFalse":          {fileName: "file.foo", expected: false},
		"GoodExtensionLowercase_ShouldReturnTrue": {fileName: "file.slf", expected: true},
		"GoodExtensionUppercase_ShouldReturnTrue": {fileName: "file.SLF", expected: true},
	}
	for name, tt := range testcases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tt.expected, (&LogAnalyzer{}).IsValidLogFileName(tt.fileName))
		})
	}
}
