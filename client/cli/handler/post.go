package handler

import (
	"io"
	"strings"

	"github.com/igorkichuk/tucows/client/cli/display"
)

type postController interface {
	GetImg(grayscale bool) (io.ReadCloser, error)
	GetQuote(key int) (string, error)
}

type imgDisplay interface {
	GetDisplayString(termWidth int, s io.Reader) (strings.Builder, error)
}

type PostHandler struct {
	l display.Logger
	c postController
	d imgDisplay
}

func NewPostHandler(c postController) PostHandler {
	return PostHandler{
		c: c,
		d: display.ImgDisplay{},
		l: display.Lgr{},
	}
}

func (h PostHandler) ShowRandomPost(grayscale bool, termWidth int, key int) {
	var builder strings.Builder
	im, err := h.c.GetImg(grayscale)
	if im != nil {
		defer im.Close()
	}
	if err != nil {
		h.l.Log(err)
	} else {
		builder, err = h.d.GetDisplayString(termWidth, im)
		if err != nil {
			h.l.Log(err)
		} else {
			h.l.Log(builder.String())
		}
	}

	qt, err := h.c.GetQuote(key)
	if err != nil {
		h.l.Log(err)
	} else {
		h.l.Log(qt)
	}
}
