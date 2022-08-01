package go_quickLogger

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, fmt.Sprintf("%-8s", "INFO:"), log.LstdFlags)
	traceLogger = log.New(os.Stdout, fmt.Sprintf("%-8s", "TRACE:"), log.LstdFlags|log.LstdFlags)
	warnLogger  = log.New(os.Stdout, fmt.Sprintf("%-8s", "WARN:"), log.LstdFlags|log.Lshortfile)
	errorLogger = log.New(os.Stdout, fmt.Sprintf("%-8s", "ERROR:"), log.LstdFlags|log.Lshortfile)
	debugLogger = log.New(os.Stdout, fmt.Sprintf("%-8s", "DEBUG:"), log.LstdFlags|log.Lshortfile)
	fatalLogger = log.New(os.Stdout, fmt.Sprintf("%-8s", "FATAL:"), log.LstdFlags|log.Lshortfile)
)

func Info(v ...interface{}) {
	infoLogger.Println(v...)
}
func Infof(format string, v ...interface{}) {
	infoLogger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	warnLogger.Println(v...)
}
func Warnf(format string, v ...interface{}) {
	warnLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	errorLogger.Println(v...)
}
func Errorf(format string, v ...interface{}) {
	errorLogger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	fatalLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	fatalLogger.Fatalf(format, v...)
}
func Trace(v ...interface{}) {
	traceLogger.Println(v...)
}
func Tracef(format string, v ...interface{}) {
	traceLogger.Printf(format, v...)
}

func Debug(v ...interface{}) {
	debugLogger.Println(v...)
}
