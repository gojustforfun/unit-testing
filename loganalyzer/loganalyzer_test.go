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
