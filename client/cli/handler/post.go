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
	imgC := make(chan interface{})
	qtC := make(chan interface{})
	go h.getImg(grayscale, termWidth, imgC)
	go h.getQuote(key, qtC)

	h.l.Log(<-imgC)
	h.l.Log(<-qtC)
}

func (h PostHandler) getImg(grayscale bool, termWidth int, resC chan<- interface{}) {
	var builder strings.Builder
	im, err := h.c.GetImg(grayscale)
	if im != nil {
		defer im.Close()
	}
	if err != nil {
		resC <- err
	} else {
		builder, err = h.d.GetDisplayString(termWidth, im)
		if err != nil {
			resC <- err
		} else {
			resC <- builder.String()
		}
	}
}

func (h PostHandler) getQuote(key int, resC chan<- interface{}) {
	qt, err := h.c.GetQuote(key)
	if err != nil {
		resC <- err
	} else {
		resC <- qt
	}
}
