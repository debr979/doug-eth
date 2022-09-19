package logging

import (
	"io"
	"log"
	"os"
)

type logging struct {
}

var Logging logging

func (r *logging) ErrorLogging(msg ...interface{}) {
	errFileName := `log.out`
	logFile, _ := os.OpenFile(errFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errLog := log.New(io.MultiWriter(logFile, os.Stderr), `[Error]`, log.Ldate|log.Ltime|log.Llongfile)
	errLog.Println(msg...)
}

func (r *logging) WarningLogging(msg ...interface{}) {
	errFileName := `log.out`
	logFile, _ := os.OpenFile(errFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errLog := log.New(io.MultiWriter(logFile, os.Stderr), `[Warning]`, log.Ldate|log.Ltime|log.Llongfile)
	errLog.Println(msg...)
}

func (r *logging) InfoLogging(msg ...interface{}) {
	errFileName := `log.out`
	logFile, _ := os.OpenFile(errFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	errLog := log.New(io.MultiWriter(logFile, os.Stderr), `[Info]`, log.Ldate|log.Ltime|log.Llongfile)
	errLog.Println(msg...)
}
