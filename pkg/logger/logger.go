package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

type Logger struct {
	verbose bool
	infoLog *log.Logger
	errLog  *log.Logger
}

func New(verbose bool) *Logger {
	return &Logger{
		verbose: verbose,
		infoLog: log.New(os.Stdout, "", 0),
		errLog:  log.New(os.Stderr, "", 0),
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf("%s: %v", msg, args)
	}
	l.infoLog.Println(color.BlueString("ℹ️  " + msg))
}

func (l *Logger) Success(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf("%s: %v", msg, args)
	}
	l.infoLog.Println(color.GreenString("✅ " + msg))
}

func (l *Logger) Warning(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf("%s: %v", msg, args)
	}
	l.infoLog.Println(color.YellowString("⚠️  " + msg))
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf("%s: %v", msg, args)
	}
	l.errLog.Println(color.RedString("❌ " + msg))
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if !l.verbose {
		return
	}
	if len(args) > 0 {
		msg = fmt.Sprintf("%s: %v", msg, args)
	}
	l.infoLog.Println(color.CyanString("🔍 " + msg))
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.Error(msg, args...)
	os.Exit(1)
}
