package handler

import (
	"fmt"
	"github.com/igorkichuk/tucows/client/web/controller"
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
}

func NewPostHandler(c controller.PostController) PostHandler {
	return PostHandler{
		c: c,
	}
}

func (h PostHandler) ShowRandomPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	grayscale, key := getPostParams(req)
	url, err := h.c.GetImg(grayscale)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "<p>img is not available</p>")
	} else {
		fmt.Fprint(w, "<img src='"+url+"'><br>")
	}

	qt, err := h.c.GetQuote(key)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "<p>...</p>")
	} else {
		fmt.Fprint(w, qt)
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
