package logger

import (
	"fmt"
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
	var (
		pc   uintptr
		fn   string
		line int
	)

	date := time.Now().Format("Mon Jan 2 15:04:05 2006")
	fmt.Printf("%s | Func: %s | File: [%s:%d] | Error: %v | Msg %s\n", date, runtime.FuncForPC(pc).Name(), fn, line, err, str)
}
