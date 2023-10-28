package handler

import (
	"github.com/igorkichuk/tucows/client/web/controller"
	"github.com/igorkichuk/tucows/client/web/display"
	"net/http"
	"strconv"

	"github.com/igorkichuk/tucows/common"
)

type postController interface {
	GetImg(grayscale bool) (string, error)
	GetQuote(key int) (string, error)
}

type PostHandler struct {
	c postController
	l common.Logger
}

func NewPostHandler(c controller.PostController) PostHandler {
	return PostHandler{
		c: c,
		l: common.DefaultLogger,
	}
}

func (h PostHandler) ShowRandomPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	grayscale, key := getPostParams(req)
	imgC := make(chan display.StringApiRes)
	qtC := make(chan display.StringApiRes)
	go h.getImg(grayscale, imgC)
	go h.getQuote(key, qtC)

	display.PrintPostRes(h.l, w, <-imgC)
	display.PrintPostRes(h.l, w, <-qtC)
}

func (h PostHandler) getImg(grayscale bool, resC chan<- display.StringApiRes) {
	url, err := h.c.GetImg(grayscale)
	if err != nil {
		resC <- display.StringApiRes{Err: err, Res: "<p>img is not available</p>"}
	} else {
		resC <- display.StringApiRes{Res: "<img src='" + url + "'><br>"}
	}
}

func (h PostHandler) getQuote(key int, resC chan<- display.StringApiRes) {
	qt, err := h.c.GetQuote(key)
	if err != nil {
		resC <- display.StringApiRes{Err: err, Res: "<p>quote is not available</p>"}
	} else {
		resC <- display.StringApiRes{Res: qt}
	}
}

func getPostParams(r *http.Request) (bool, int) {
	keyS := r.URL.Query().Get("key")
	key, err := strconv.Atoi(keyS)
	if err != nil {
		key = common.DefaultKey
	}
	greyS := r.URL.Query().Get("grayscale")
	gray, err := strconv.ParseBool(greyS)
	if err != nil {
		gray = common.DefaultGrayscale
	}

	return gray, key
}
