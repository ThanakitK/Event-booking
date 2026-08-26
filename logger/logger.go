// logger/logger.go
package logger

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

// GetLogger returns a configured logrus logger with component and function info
func GetLogger() *log.Entry {
	pc, file, _, _ := runtime.Caller(1) // Get caller info one level up
	functionName := strings.Split(runtime.FuncForPC(pc).Name(), ".")[len(strings.Split(runtime.FuncForPC(pc).Name(), "."))-1]
	fileName := strings.Split(file, "/")[len(strings.Split(file, "/"))-1]

	return log.WithFields(log.Fields{
		"component": fileName,
		"function":  functionName,
	})
}
