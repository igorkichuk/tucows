package handler

import (
	"github.com/igorkichuk/tucows/client/cli/controller"
	"github.com/igorkichuk/tucows/client/cli/display"
	"io"
	"strings"
	"testing"
	"time"
)

func BenchmarkShowRandPost(b *testing.B) {
	c := controller.PostMock{
		GetImgFunc: func(grayscale bool) (io.ReadCloser, error) {
			time.Sleep(time.Second)
			return nil, nil
		},
		GetQuoteFunc: func(key int) (string, error) {
			time.Sleep(500 * time.Millisecond)
			return "some quote", nil
		},
	}
	d := display.PostMock{GetDisplayStringFunc: func(termWidth int, s io.Reader) (strings.Builder, error) {
		bil := strings.Builder{}
		bil.Write([]byte("some terminal output"))
		return bil, nil
	}}

	h := PostHandler{
		c: c,
		d: d,
		l: display.Lgr{},
	}

	b.ResetTimer()
	h.ShowRandomPost(false, 400, 3)
}
