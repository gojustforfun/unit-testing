package loganalyzer

import "strings"

type LogAnalyzer struct{}

func (l *LogAnalyzer) IsValidLogFileName(fileName string) bool {
	if !strings.HasSuffix(fileName, ".SLF") {
		return false
	}
	return true
}
