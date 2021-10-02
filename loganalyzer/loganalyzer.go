package loganalyzer

import "strings"

type LogAnalyzer struct{}

func (l *LogAnalyzer) IsValidLogFileName(fileName string) bool {
	if !strings.HasSuffix(strings.ToLower(fileName), ".slf") {
		return false
	}
	return true
}
