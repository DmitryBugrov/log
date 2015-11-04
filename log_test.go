package log

import "testing"

func TestLoglevelTrace(t *testing.T) {
	Init(LogLevelTrace)
	Print(LogLevelTrace, "Test trace message")
	Print(LogLevelError, "Test error message")

}

func TestLoglevelError(t *testing.T) {
	Init(LogLevelError)
	Print(LogLevelTrace, "Test trace message")
	Print(LogLevelError, "Test error message")

}
