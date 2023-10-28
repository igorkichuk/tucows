package common

import (
	"fmt"
	"io"
)

var DefaultLogger = Lgr{}

type Logger interface {
	Logln(a ...any)
	Flog(w io.Writer, a ...any) (n int, err error)
}

func NewLogger() Logger {
	return Lgr{}
}

type Lgr struct{}

func (Lgr) Logln(a ...any) {
	fmt.Println(a...)
}

func (Lgr) Flog(w io.Writer, a ...any) (n int, err error) {
	return fmt.Fprint(w, a...)
}
