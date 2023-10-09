package quote

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestQT_Get(t *testing.T) {
	tests := map[string]struct {
		returnStatus int
		returnBody   []byte
		returnErr    error
		exp          string
		expErr       bool
	}{
		"correct": {
			returnStatus: http.StatusOK,
			returnBody:   []byte(`body`),
			returnErr:    nil,
			exp:          "body",
			expErr:       false,
		},
		"wrongStatus": {
			returnStatus: http.StatusNotFound,
			returnBody:   []byte(`body`),
			returnErr:    nil,
			exp:          "",
			expErr:       true,
		},
		"errRequest": {
			returnStatus: http.StatusBadGateway,
			returnBody:   []byte(`body`),
			returnErr:    fmt.Errorf("broken connection"),
			exp:          "",
			expErr:       true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			httpmock.RegisterResponder("GET", "http://api.forismatic.com/api/1.0/",
				func(req *http.Request) (*http.Response, error) {
					resp := httpmock.NewBytesResponse(tt.returnStatus, tt.returnBody)
					return resp, tt.returnErr
				},
			)

			q := QT{
				client:     http.DefaultClient,
				respFormat: TextFormat,
			}

			got, err := q.Get(120)
			if (err != nil) != tt.expErr {
				t.Errorf("Get() error = %v, expErr %v", err, tt.expErr)
				return
			}
			if got != tt.exp {
				t.Errorf("Get() got = %v, want %v", got, tt.exp)
			}
		})
	}
}
