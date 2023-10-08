package display

import "fmt"

type Logger interface {
	Log(a ...any)
}

type Lgr struct{}

func (Lgr) Log(a ...any) {
	fmt.Println(a...)
}
