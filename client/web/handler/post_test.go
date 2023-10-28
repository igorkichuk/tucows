package handler

import (
	"errors"
	"github.com/igorkichuk/tucows/client/web/controller"
	"github.com/igorkichuk/tucows/common"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler_ShowRandomPost(t *testing.T) {
	tests := map[string]struct {
		c       postController
		params  string
		expLog  string
		expResp string
	}{
		"correct_common_params": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (string, error) {
					if grayscale != common.DefaultGrayscale {
						t.Errorf("expected grayscale is %t, got %t", common.DefaultGrayscale, grayscale)
					}
					return "some img", nil
				},
				GetQuoteFunc: func(key int) (string, error) {
					if key != common.DefaultKey {
						t.Errorf("expected key is %d, got %d", common.DefaultKey, key)
					}
					return "<p>Age does not protect you from love. But love, to some extent, protects you from age. (Anais Nin)</p>", nil
				},
			},
			expLog:  "",
			expResp: "<img src='some img'><br><p>Age does not protect you from love. But love, to some extent, protects you from age. (Anais Nin)</p>",
		},
		"correct_custom_params": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (string, error) {
					if grayscale != true {
						t.Errorf("expected grayscale is %t, got %t", true, grayscale)
					}
					return "some img", nil
				},
				GetQuoteFunc: func(key int) (string, error) {
					if key != 5 {
						t.Errorf("expected key is %d, got %d", 5, key)
					}
					return "<p>Age does not protect you from love. But love, to some extent, protects you from age. (Anais Nin)</p>", nil
				},
			},
			params:  "grayscale=true&key=5",
			expLog:  "",
			expResp: "<img src='some img'><br><p>Age does not protect you from love. But love, to some extent, protects you from age. (Anais Nin)</p>",
		},
		"get_img_err": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (string, error) {
					return "", errors.New("some error")
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "<p>Age does not protect you from love. But love, to some extent, protects you from age. (Anais Nin)</p>", nil
				},
			},
			expLog:  "some error\n",
			expResp: "<p>img is not available</p><p>Age does not protect you from love. But love, to some extent, protects you from age. (Anais Nin)</p>",
		},
		"get_qt_err": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (string, error) {
					return "some img", nil
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "", errors.New("some error")
				},
			},
			expLog:  "some error\n",
			expResp: "<img src='some img'><br><p>quote is not available</p>",
		},
		"two_api_err": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (string, error) {
					return "", errors.New("some error")
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "", errors.New("some another error")
				},
			},
			expLog:  "some error\nsome another error\n",
			expResp: "<p>img is not available</p><p>quote is not available</p>",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := common.NewLoggerMock()
			h := PostHandler{
				c: tt.c,
				l: l,
			}
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/post?"+tt.params, nil)

			h.ShowRandomPost(w, r)
			s := l.GetLog()
			if s != tt.expLog {
				t.Errorf("ShowRandPost() log exp: %s; got: %s", tt.expLog, s)
			}
			resp := w.Body.String()
			if resp != tt.expResp {
				t.Errorf("ShowRandPost() resp exp: %s; got: %s", tt.expResp, resp)
			}
		})
	}
}
