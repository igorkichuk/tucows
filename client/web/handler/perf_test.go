package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/igorkichuk/tucows/client/web/controller"
	"github.com/igorkichuk/tucows/common"
)

func BenchmarkShowRandPost(b *testing.B) {
	c := controller.PostMock{
		GetImgFunc: func(grayscale bool) (string, error) {
			time.Sleep(time.Second)
			return "some img", nil
		},
		GetQuoteFunc: func(key int) (string, error) {
			time.Sleep(500 * time.Millisecond)
			return "some quote", nil
		},
	}

	h := PostHandler{
		c: c,
		l: common.DefaultLogger,
	}

	b.ResetTimer()
	h.ShowRandomPost(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/post", nil))
}
