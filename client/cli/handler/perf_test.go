package handler

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/igorkichuk/tucows/client/cli/controller"
	"github.com/igorkichuk/tucows/client/cli/display"
	"github.com/igorkichuk/tucows/common"
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
		l: common.DefaultLogger,
	}

	b.ResetTimer()
	h.ShowRandomPost(false, 400, 3)
}
