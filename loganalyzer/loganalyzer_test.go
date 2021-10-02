package loganalyzer_test

import (
	"testing"

	. "github.com/gojustforfun/unit-testing/loganalyzer"
	"github.com/stretchr/testify/assert"
)

func TestIsValidLogFileName_BadExtension_ReturnFalse(t *testing.T) {
	analyzer := LogAnalyzer{}
	result := analyzer.IsValidLogFileName("file.foo")
	assert.Equal(t, false, result)
}

func TestIsValidLogFileName_GoodExtensionLowercase_ReturnTrue(t *testing.T) {
	analyzer := LogAnalyzer{}
	result := analyzer.IsValidLogFileName("file.slf")
	assert.Equal(t, true, result)
}

func TestIsValidLogFileName_GoodExtensionUppercase_ReturnTrue(t *testing.T) {
	analyzer := LogAnalyzer{}
	result := analyzer.IsValidLogFileName("file.SLF")
	assert.Equal(t, true, result)
}
