package display

import (
	"io"
	"strings"
)

type PostMock struct {
	GetDisplayStringFunc func(termWidth int, s io.Reader) (strings.Builder, error)
}

func (m PostMock) GetDisplayString(termWidth int, s io.Reader) (strings.Builder, error) {
	return m.GetDisplayStringFunc(termWidth, s)
}
