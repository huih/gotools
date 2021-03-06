package logs

import (
	"testing"
)

func TestLogSetLevel(t *testing.T){
	SetLevel(LOG_INFO)
}

func TestLogWriteToFile(t *testing.T) {
	SetFilePath("D:\\work\\log.txt")
	Info("hello web spider log file")
	Info("hello web spider log file")
}

func TestLogWriteToStdOut(t *testing.T) {
	Info("write log to stdout")
}

func BenchmarkLogInfo(b *testing.B) {
	Info("write log to stdout")
}

func TestLogSWriteToStdOut(t *testing.T) {
	InfoS("write log to stdout")
}
