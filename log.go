package log

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	LogLevelTrace = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

type Log struct {
	logLevel  int
	funcName  bool
	line      bool
	timeStamp bool
}

//Init Log
//Init(LogLevel Log,funcName bool,line bool,timeStamp boot)
func (l *Log) Init(ll int, param ...bool) bool {
	//	logLevel = ll
	switch len(param) {
	case 0:
		l.funcName = true
		l.line = true
		l.timeStamp = true
	case 1:
		l.funcName = param[0]
		l.line = true
		l.timeStamp = true
	case 2:
		l.funcName = param[0]
		l.line = param[1]
		l.timeStamp = true
	case 3:
		l.funcName = param[0]
		l.line = param[1]
		l.timeStamp = param[2]
	default:
		return false
	}
	return true
}

//Print log message
func (l *Log) Print(ll int, msg ...interface{}) string {
	pc, _, num_line, _ := runtime.Caller(1)

	var tmestamp, funcNameStr, lineStr string
	if l.timeStamp {
		tmestamp = time.Now().Format(time.StampMilli)
	}
	if l.funcName {
		funcNameStr = runtime.FuncForPC(pc).Name()
	}
	if l.line {
		lineStr = ":" + strconv.Itoa(num_line)
	}

	result := ""
	if ll >= l.logLevel {
		switch ll {
		case LogLevelError:
			result = tmestamp + " [Error] " + funcNameStr + lineStr + strings.Join(interfaceToString(msg)[:], "")
		case LogLevelWarning:
			result = tmestamp + " [Warning] " + funcNameStr + lineStr + strings.Join(interfaceToString(msg)[:], "")
		case LogLevelInfo:
			result = tmestamp + " [Info] " + funcNameStr + lineStr + strings.Join(interfaceToString(msg)[:], "")
		case LogLevelTrace:
			result = tmestamp + " [Trace] " + funcNameStr + lineStr + strings.Join(interfaceToString(msg)[:], "")
		}

	}
	if result != "" {
		fmtPrint(result)
	}
	return result

}

func interfaceToString(a []interface{}) []string {
	result := make([]string, len(a))
	for i := range a {
		result[i] = " " + a[i].(string)
	}
	return result
}

func fmtPrint(msg string) {
	fmt.Println(msg)
}
