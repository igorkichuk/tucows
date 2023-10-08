package display

import (
	"bytes"
	"fmt"
)

func NewLoggerMock() LoggerMock {
	return LoggerMock{buf: bytes.NewBuffer([]byte(""))}
}

type LoggerMock struct {
	buf *bytes.Buffer
}

func (l LoggerMock) Log(a ...any) {
	l.buf.Write([]byte(fmt.Sprint(a...) + "\n"))
}

func (l LoggerMock) GetLog() string {
	return l.buf.String()
}
