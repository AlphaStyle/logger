package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// Info will print out usefull information
func Info(str string) {
	date := time.Now().Format("Mon Jan 2 15:04:05 2006")
	fmt.Printf("%s | %s\n", date, str)
}

// Infof is like Info but with format
func Infof(format string, a ...interface{}) {
	date := time.Now().Format("Mon Jan 2 15:04:05 2006")
	fmt.Printf(date+" | "+format, a...)
}

// Error is to log errors
func Error(err error, str string) {
	pc, fn, line, _ := runtime.Caller(1)

	date := time.Now().Format("Mon Jan 2 15:04:05 2006")
	fmt.Printf("%s\nFunc: %s\nFile: [%s:%d]\nError: %v\nMsg: %s\n", date, runtime.FuncForPC(pc).Name(), fn, line, err, str)
	os.Exit(2)
}
