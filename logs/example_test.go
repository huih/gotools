package logs_test

import (
	"github.com/gotools/logs"
)

func ExampleLog(){
	logs.SetFilePath("D:\\work\\log.txt")
	logs.Info("hello web spider log file: %s", "info")
	logs.Debug("hello web spider log file: %s", "debug")
	logs.Warning("hello web spider log file: %s", "warning")
	logs.Error("hello web spider log file: %s", "error")
	logs.Fatal("hello web spider log file: %s", "fatal")
}