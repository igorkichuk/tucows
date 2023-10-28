package handler

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/igorkichuk/tucows/client/cli/controller"
	"github.com/igorkichuk/tucows/client/cli/display"
	"github.com/igorkichuk/tucows/common"
)

func TestPostHandler_ShowRandomPost(t *testing.T) {
	tests := map[string]struct {
		c   postController
		d   imgDisplay
		exp string
	}{
		"correct": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (io.ReadCloser, error) {
					return nil, nil
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "some quote", nil
				},
			},
			d: display.PostMock{GetDisplayStringFunc: func(termWidth int, s io.Reader) (strings.Builder, error) {
				b := strings.Builder{}
				b.Write([]byte("some terminal output"))
				return b, nil
			}},
			exp: "some terminal output\nsome quote\n",
		},
		"getImgFailed": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (io.ReadCloser, error) {
					return nil, fmt.Errorf("get image failed")
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "some quote", nil
				},
			},
			d: display.PostMock{GetDisplayStringFunc: func(termWidth int, s io.Reader) (strings.Builder, error) {
				b := strings.Builder{}
				b.Write([]byte("some terminal output"))
				return b, nil
			}},
			exp: "get image failed\nsome quote\n",
		},
		"getQuoteFailed": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (io.ReadCloser, error) {
					return nil, nil
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "", fmt.Errorf("get quote failed")
				},
			},
			d: display.PostMock{GetDisplayStringFunc: func(termWidth int, s io.Reader) (strings.Builder, error) {
				b := strings.Builder{}
				b.Write([]byte("some terminal output"))
				return b, nil
			}},
			exp: "some terminal output\nget quote failed\n",
		},
		"bothFailed": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (io.ReadCloser, error) {
					return nil, fmt.Errorf("get image failed")
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "", fmt.Errorf("get quote failed")
				},
			},
			d: display.PostMock{GetDisplayStringFunc: func(termWidth int, s io.Reader) (strings.Builder, error) {
				b := strings.Builder{}
				b.Write([]byte("some terminal output"))
				return b, nil
			}},
			exp: "get image failed\nget quote failed\n",
		},
		"displayFailed": {
			c: controller.PostMock{
				GetImgFunc: func(grayscale bool) (io.ReadCloser, error) {
					return nil, nil
				},
				GetQuoteFunc: func(key int) (string, error) {
					return "some quote", nil
				},
			},
			d: display.PostMock{GetDisplayStringFunc: func(termWidth int, s io.Reader) (strings.Builder, error) {
				return strings.Builder{}, fmt.Errorf("display string failed")
			}},
			exp: "display string failed\nsome quote\n",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			l := common.NewLoggerMock()
			h := PostHandler{
				c: tt.c,
				d: tt.d,
				l: l,
			}
			h.ShowRandomPost(true, 120, 10)
			s := l.GetLog()
			if s != tt.exp {
				t.Errorf("ShowRandomPost() exp: %s; got: %s", tt.exp, s)
			}
		})
	}
}
