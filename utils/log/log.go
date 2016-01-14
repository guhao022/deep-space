package log

import "github.com/num5/log5"

var l *log5.Log

func Infof(format string, v ...interface{}) {
	l.Infof(format, v)
}
func Info(v ...interface{}) {
	l.Info(v)
}

func Tracf(format string, v ...interface{}) {
	l.Tracf(format, v)
}
func Trac(v ...interface{}) {
	l.Trac(v)
}

func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v)
}
func Warn(v ...interface{}) {
	l.Warn(v)
}

func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v)
}
func Error(v ...interface{}) {
	l.Error(v)
}

func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v)
}
func Fatal(v ...interface{}) {
	l.Fatal(v)
}


func init() {
	l = log5.NewLog(1000)
	//l.SetFuncCall(true)
	//l.SetFuncCallDepth(3)
}