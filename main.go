package go_quickLogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LoggersStruct struct {
	InfoLogger  *log.Logger
	TraceLogger *log.Logger
	WarnLogger  *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
	FatalLogger *log.Logger
}

var (
	Loggers = LoggersStruct{
		InfoLogger:  log.New(os.Stdout, fmt.Sprintf("%-8s", "INFO:"), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix),
		TraceLogger: log.New(os.Stdout, fmt.Sprintf("%-8s", "TRACE:"), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix),
		WarnLogger:  log.New(os.Stdout, fmt.Sprintf("%-8s", "WARN:"), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix),
		ErrorLogger: log.New(os.Stdout, fmt.Sprintf("%-8s", "ERROR:"), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix),
		DebugLogger: log.New(os.Stdout, fmt.Sprintf("%-8s", "DEBUG:"), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix),
		FatalLogger: log.New(os.Stdout, fmt.Sprintf("%-8s", "FATAL:"), log.LstdFlags|log.Lmicroseconds|log.Lmsgprefix),
	}
	Tracing   = false
	LogToFile = false
)

func getLogFileName() string {
	return "logs/" + time.Now().Format("2006-01-02") + ".log"
}

func Info(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.InfoLogger.SetOutput(f)
		Loggers.InfoLogger.Println(v...)
		f.Close()
		Loggers.InfoLogger.SetOutput(os.Stdout)
	}
	Loggers.InfoLogger.Println(v...)
}
func Infof(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.InfoLogger.SetOutput(f)
		Loggers.InfoLogger.Printf(format, v...)
		f.Close()
		Loggers.InfoLogger.SetOutput(os.Stdout)
	}
	Loggers.InfoLogger.Printf(format, v...)
}

func Warn(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.WarnLogger.SetOutput(f)
		Loggers.WarnLogger.Println(v...)
		f.Close()
		Loggers.WarnLogger.SetOutput(os.Stdout)
	}
	Loggers.WarnLogger.Println(v...)
}
func Warnf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.WarnLogger.SetOutput(f)
		Loggers.WarnLogger.Printf(format, v...)
		f.Close()
		Loggers.WarnLogger.SetOutput(os.Stdout)
	}
	Loggers.WarnLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.ErrorLogger.SetOutput(f)
		Loggers.ErrorLogger.Println(v...)
		f.Close()
		Loggers.ErrorLogger.SetOutput(os.Stdout)
	}
	Loggers.ErrorLogger.Println(v...)
}
func Errorf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.ErrorLogger.SetOutput(f)
		Loggers.ErrorLogger.Printf(format, v...)
		f.Close()
		Loggers.ErrorLogger.SetOutput(os.Stdout)
	}
	Loggers.ErrorLogger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.FatalLogger.SetOutput(f)
		Loggers.FatalLogger.Println(v...)
		f.Close()
		Loggers.FatalLogger.SetOutput(os.Stdout)
	}
	Loggers.FatalLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.FatalLogger.SetOutput(f)
		Loggers.FatalLogger.Printf(format, v...)
		f.Close()
		Loggers.FatalLogger.SetOutput(os.Stdout)
	}
	Loggers.FatalLogger.Fatalf(format, v...)
}
func Trace(v ...interface{}) {
	if Tracing {
		if LogToFile {
			f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			Loggers.TraceLogger.SetOutput(f)
			Loggers.TraceLogger.Println(v...)
			f.Close()
			Loggers.TraceLogger.SetOutput(os.Stdout)
		}
		Loggers.TraceLogger.Println(v...)
	}
}
func Tracef(format string, v ...interface{}) {
	if Tracing {
		if LogToFile {
			f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			Loggers.TraceLogger.SetOutput(f)
			Loggers.TraceLogger.Printf(format, v...)
			f.Close()
			Loggers.TraceLogger.SetOutput(os.Stdout)
		}
		Loggers.TraceLogger.Printf(format, v...)
	}
}

func Debug(v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.DebugLogger.SetOutput(f)
		Loggers.DebugLogger.Println(v...)
		f.Close()
		Loggers.DebugLogger.SetOutput(os.Stdout)
	}
	Loggers.DebugLogger.Println(v...)
}

func Debugf(format string, v ...interface{}) {
	if LogToFile {
		f, _ := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		Loggers.DebugLogger.SetOutput(f)
		Loggers.DebugLogger.Printf(format, v...)
		f.Close()
		Loggers.DebugLogger.SetOutput(os.Stdout)
	}
	Loggers.DebugLogger.Printf(format, v...)
}
