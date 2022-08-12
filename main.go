package go_quickLogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	InfoLogger  = log.New(os.Stdout, fmt.Sprintf("%s %-7s", time.Now().Format("2006-01-02 15:04:05.000"), "INFO:"), 0)
	TraceLogger = log.New(os.Stdout, fmt.Sprintf("%s %-7s", time.Now().Format("2006-01-02 15:04:05.000"), "TRACE:"), 0)
	WarnLogger  = log.New(os.Stdout, fmt.Sprintf("%s %-7s", time.Now().Format("2006-01-02 15:04:05.000"), "WARN:"), 0)
	ErrorLogger = log.New(os.Stdout, fmt.Sprintf("%s %-7s", time.Now().Format("2006-01-02 15:04:05.000"), "ERROR:"), 0)
	DebugLogger = log.New(os.Stdout, fmt.Sprintf("%s %-7s", time.Now().Format("2006-01-02 15:04:05.000"), "DEBUG:"), 0)
	FatalLogger = log.New(os.Stdout, fmt.Sprintf("%s %-7s", time.Now().Format("2006-01-02 15:04:05.000"), "FATAL:"), 0)
	Tracing     = false
	LogToFile   = false
)

func getLogFileName() string {
	return "logs/" + time.Now().Format("2006-01-02") + ".log"
}

func Info(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		InfoLogger.SetOutput(f)
		InfoLogger.Println(v...)
		f.Close()
		InfoLogger.SetOutput(os.Stdout)
	}
	InfoLogger.Println(v...)
}
func Infof(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		InfoLogger.SetOutput(f)
		InfoLogger.Printf(format, v...)
		f.Close()
		InfoLogger.SetOutput(os.Stdout)
	}
	InfoLogger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		WarnLogger.SetOutput(f)
		WarnLogger.Println(v...)
		f.Close()
		WarnLogger.SetOutput(os.Stdout)
	}
	WarnLogger.Println(v...)
}
func Warnf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		WarnLogger.SetOutput(f)
		WarnLogger.Printf(format, v...)
		f.Close()
		WarnLogger.SetOutput(os.Stdout)
	}
	WarnLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		ErrorLogger.SetOutput(f)
		ErrorLogger.Println(v...)
		f.Close()
		ErrorLogger.SetOutput(os.Stdout)
	}
	ErrorLogger.Println(v...)
}
func Errorf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		ErrorLogger.SetOutput(f)
		ErrorLogger.Printf(format, v...)
		f.Close()
		ErrorLogger.SetOutput(os.Stdout)
	}
	ErrorLogger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		FatalLogger.SetOutput(f)
		FatalLogger.Println(v...)
		f.Close()
		FatalLogger.SetOutput(os.Stdout)
	}
	FatalLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		FatalLogger.SetOutput(f)
		FatalLogger.Printf(format, v...)
		f.Close()
		FatalLogger.SetOutput(os.Stdout)
	}
	FatalLogger.Fatalf(format, v...)
}
func Trace(v ...interface{}) {
	if Tracing {
		if LogToFile {
			f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			TraceLogger.SetOutput(f)
			TraceLogger.Println(v...)
			f.Close()
			TraceLogger.SetOutput(os.Stdout)
		}
		TraceLogger.Println(v...)
	}
}
func Tracef(format string, v ...interface{}) {
	if Tracing {
		if LogToFile {
			f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			TraceLogger.SetOutput(f)
			TraceLogger.Printf(format, v...)
			f.Close()
			TraceLogger.SetOutput(os.Stdout)
		}
		TraceLogger.Printf(format, v...)
	}
}

func Debug(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		DebugLogger.SetOutput(f)
		DebugLogger.Println(v...)
		f.Close()
		DebugLogger.SetOutput(os.Stdout)
	}
	DebugLogger.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		DebugLogger.SetOutput(f)
		DebugLogger.Printf(format, v...)
		f.Close()
		DebugLogger.SetOutput(os.Stdout)
	}
	DebugLogger.Printf(format, v...)
}
