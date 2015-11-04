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
)

//Init Log
func Init(ll int) {
	logLevel = ll

}

//Print log message
func Print(ll int, msg ...string) {
	pc, _, line, _ := runtime.Caller(1)
	func_and_line := runtime.FuncForPC(pc).Name() + ":" + strconv.Itoa(line)
	tmestamp := time.Now().Format(time.StampMilli)
	if ll >= logLevel {
		switch ll {
		case LogLevelError:
			fmt.Println(tmestamp, "[Error]", func_and_line, msg)
		case LogLevelWarning:
			fmt.Println(tmestamp, "[Warning]", func_and_line, msg)
		case LogLevelInfo:
			fmt.Println(tmestamp, "[Info]", func_and_line, msg)
		case LogLevelTrace:
			fmt.Println(tmestamp, "[Trace]", func_and_line, msg)
		}

	}

}
