package go_quickLogger

import (
	"fmt"
	"log"
	"os"
)

var (
	InfoLogger  = log.New(os.Stdout, fmt.Sprintf("%-7s", "INFO:"), log.LstdFlags)
	TraceLogger = log.New(os.Stdout, fmt.Sprintf("%-7s", "TRACE:"), log.LstdFlags|log.Llongfile)
	WarnLogger  = log.New(os.Stdout, fmt.Sprintf("%-7s", "WARN:"), log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(os.Stdout, fmt.Sprintf("%-7s", "ERROR:"), log.LstdFlags|log.Lshortfile)
	DebugLogger = log.New(os.Stdout, fmt.Sprintf("%-7s", "DEBUG:"), log.LstdFlags|log.Lshortfile)
	FatalLogger = log.New(os.Stdout, fmt.Sprintf("%-7s", "FATAL:"), log.LstdFlags|log.Lshortfile)
	Tracing     = false
)

func Info(v ...interface{}) {
	InfoLogger.Println(v...)
}
func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	WarnLogger.Println(v...)
}
func Warnf(format string, v ...interface{}) {
	WarnLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	ErrorLogger.Println(v...)
}
func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	FatalLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	FatalLogger.Fatalf(format, v...)
}
func Trace(v ...interface{}) {
	if Tracing {
		TraceLogger.Println(v...)
	}
}
func Tracef(format string, v ...interface{}) {
	if Tracing {
		TraceLogger.Printf(format, v...)
	}
}

func Debug(v ...interface{}) {
	DebugLogger.Println(v...)
}
