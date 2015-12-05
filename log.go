package log

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

const (
	LogLevelTrace = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

var (
	logLevel int = 3
	funcName = true
	line =true
	timeStamp=true
)

//Init Log
func Init(ll int,_timeStamp bool,_funcName bool,_line bool) {
	logLevel = ll
	funcName =_funcName
	timeStamp=_timeStamp
	line=_line
}

//Print log message
func Print(ll int, msg ...interface{}) {
	pc, _, num_line, _ := runtime.Caller(1)
	//func_and_line := runtime.FuncForPC(pc).Name() + ":" + strconv.Itoa(line)
	var tmestamp,funcNameStr,lineStr string
	if timeStamp {
		tmestamp = time.Now().Format(time.StampMilli)	
	}
	if funcName {
		funcNameStr=runtime.FuncForPC(pc).Name()
	}
	if line {
		lineStr=":" + strconv.Itoa(num_line)
	}
	
	if ll >= logLevel {
		switch ll {
		case LogLevelError:
			fmt.Println(tmestamp, "[Error]", funcNameStr,lineStr, msg)
		case LogLevelWarning:
			fmt.Println(tmestamp, "[Warning]", funcNameStr,lineStr, msg)
		case LogLevelInfo:
			fmt.Println(tmestamp, "[Info]", funcNameStr,lineStr, msg)
		case LogLevelTrace:
			fmt.Println(tmestamp, "[Trace]", funcNameStr,lineStr, msg)
		}

	}

}
