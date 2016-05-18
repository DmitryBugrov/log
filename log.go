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

//var (
//	logLevel int = 3
//	funcName = true
//	line =true
//	timeStamp=true
//)

type Log struct {
	logLevel  int
	funcName  bool
	line      bool
	timeStamp bool
}

//Init Log
func (self *Log) Init(ll int, param ...bool) {
	//	logLevel = ll
	switch len(param) {
	case 0:
		self.funcName = true
		self.line = true
		self.timeStamp = true
	case 1:
		self.funcName = param[0]
		self.line = true
		self.timeStamp = true
	case 2:
		self.funcName = param[0]
		self.line = param[1]
		self.timeStamp = true
	case 3:
		self.funcName = param[0]
		self.line = param[1]
		self.timeStamp = param[2]
	}
	//	if len(param) >0 nil {
	//		funcName = param[0]
	//	} else {
	//		funcName = true
	//	}
	//	if param[1] != nil {
	//		line = param[1]
	//	} else {
	//		line = true
	//	}
	//	if param[2] != nil {
	//		timeStamp = param[2]
	//	} else {
	//		timeStamp = true
	//	}

}

//Print log message
func (self *Log) Print(ll int, msg ...interface{}) {
	pc, _, num_line, _ := runtime.Caller(1)
	//func_and_line := runtime.FuncForPC(pc).Name() + ":" + strconv.Itoa(line)
	var tmestamp, funcNameStr, lineStr string
	if self.timeStamp {
		tmestamp = time.Now().Format(time.StampMilli)
	}
	if self.funcName {
		funcNameStr = runtime.FuncForPC(pc).Name()
	}
	if self.line {
		lineStr = ":" + strconv.Itoa(num_line)
	}

	if ll >= self.logLevel {
		switch ll {
		case LogLevelError:
			fmt.Println(tmestamp, "[Error]", funcNameStr, lineStr, msg)
		case LogLevelWarning:
			fmt.Println(tmestamp, "[Warning]", funcNameStr, lineStr, msg)
		case LogLevelInfo:
			fmt.Println(tmestamp, "[Info]", funcNameStr, lineStr, msg)
		case LogLevelTrace:
			fmt.Println(tmestamp, "[Trace]", funcNameStr, lineStr, msg)
		}

	}

}
