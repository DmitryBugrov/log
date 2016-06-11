package log

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	log *Log
)

func TestLogModule(t *testing.T) {

	Convey("init log module", t, func() {
		log = new(Log)
		So(log.Init(LogLevelError, true, false, false, false), ShouldBeFalse)
		So(log.Init(LogLevelError, true, false, false), ShouldBeTrue)

		Convey("log message with LogLevel Error", func() {
			So(log.Print(LogLevelError, "Test Error message"), ShouldEqual, " [Error] github.com/DmitryBugrov/log.TestLogModule.func1.1 Test Error message")
		})

		Convey("log message with LogLevel Warning", func() {
			So(log.Print(LogLevelWarning, "Test Warning message"), ShouldEqual, " [Warning] github.com/DmitryBugrov/log.TestLogModule.func1.2 Test Warning message")
		})

		Convey("log message with LogLevel Info", func() {
			So(log.Print(LogLevelInfo, "Test Info message"), ShouldEqual, " [Info] github.com/DmitryBugrov/log.TestLogModule.func1.3 Test Info message")
		})

		Convey("log message with LogLevel Trace", func() {
			So(log.Print(LogLevelTrace, "Test Trace message"), ShouldEqual, " [Trace] github.com/DmitryBugrov/log.TestLogModule.func1.4 Test Trace message")
		})

	})

}
