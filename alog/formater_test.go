package alog_test

import (
	"testing"
	"alog"
	"time"
)

func TestFormater(t *testing.T) {
	var formater = alog.NewBasicFormater()
	var log alog.LogParam

	tm, err := time.Parse(time.StampNano, "Jan 02 15:04:05.123456789")
	if err != nil {
		t.Error("formater test error, parse datetime failed, err=" + err.Error())
	}

	log.Time = tm
	log.Level = 0
	log.Args = make([]interface{}, 0)
	log.Args = append(log.Args, "Hello")
	log.Args = append(log.Args, 1)
	log.Args = append(log.Args, "world")

	var dst = "15:04:05.123 [Debug ] Hello 1 world"

	out := formater.Format(&log)

	if dst != out {
		t.Error("formater test error, dst=\"" + dst + "\", out=\"" + out + "\"")
	}
}
