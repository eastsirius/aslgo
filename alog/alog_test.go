package alog_test

import "testing"
import "alog"

func TestLog(t *testing.T) {
	alog.NotifyPrint("Hello log world!")
}
