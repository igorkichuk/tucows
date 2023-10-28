package common

import (
	"bytes"
	"fmt"
	"io"
)

func NewLoggerMock() LoggerMock {
	return LoggerMock{buf: bytes.NewBuffer([]byte(""))}
}

type LoggerMock struct {
	buf *bytes.Buffer
}

func (l LoggerMock) Logln(a ...any) {
	l.buf.Write([]byte(fmt.Sprint(a...) + "\n"))
}

func (l LoggerMock) Flog(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprint(w, a...)
}

func (l LoggerMock) GetLog() string {
	return l.buf.String()
}
