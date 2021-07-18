package Log

import (
	"log"
	"os"
	"sync"
)

type defaultLogger struct{
	infoLogger    *log.Logger
	warnLogger *log.Logger
	errorLogger   *log.Logger
}

var current ILogger
var once sync.Once

const flag = log.Ldate | log.Lmicroseconds

func Current() ILogger {
	once.Do(func(){
		defaultLogger := defaultLogger{
			infoLogger:  log.New(os.Stdout, "INFO: ", flag),
			warnLogger:  log.New(os.Stdout, "WARN: ", flag),
			errorLogger: log.New(os.Stdout, "ERROR: ", flag),
		}
		current = &defaultLogger
	})
	return current
}

func (defaultLogger *defaultLogger) LogInfo(message string){
	defaultLogger.infoLogger.Println(message)
}

func (defaultLogger *defaultLogger) LogWarn(message string){
	defaultLogger.warnLogger.Println(message)
}

func (defaultLogger *defaultLogger) LogError(err error){
	defaultLogger.errorLogger.Println(err.Error())
}