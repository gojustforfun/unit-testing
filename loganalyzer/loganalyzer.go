package loganalyzer

import "strings"

type LogAnalyzer struct{}

func (l *LogAnalyzer) IsValidLogFileName(fileName string) bool {
	return strings.HasSuffix(strings.ToLower(fileName), ".slf")
}
