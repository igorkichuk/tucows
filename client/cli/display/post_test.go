package display

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_getDisplayString(t *testing.T) {
	fCorrect, err := os.ReadFile("./tests/img220x220")
	if err != nil {
		t.Error(err)
	}
	answCorrrect, err := os.ReadFile("./tests/term_disp")
	if err != nil {
		t.Error(err)
	}
	sameWidth, err := os.ReadFile("./tests/same_width")
	if err != nil {
		t.Error(err)
	}

	tests := map[string]struct {
		termWidth int
		s         io.Reader
		exp       string
		expErr    bool
	}{
		"correct": {
			termWidth: 50,
			s:         bytes.NewBuffer(fCorrect),
			exp:       string(answCorrrect),
			expErr:    false,
		},
		"sameTermWidth": {
			termWidth: 220,
			s:         bytes.NewBuffer(fCorrect),
			exp:       string(sameWidth),
			expErr:    false,
		},
		"wrongImage": {
			termWidth: 220,
			s:         strings.NewReader("wrong image"),
			expErr:    true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := ImgDisplay{}.GetDisplayString(tt.termWidth, tt.s)
			if (err != nil) != tt.expErr {
				t.Errorf("GetDisplayString() error = %v, wantErr %v", err, tt.expErr)
				return
			}
			if got.String() != tt.exp {
				t.Errorf("GetDisplayString() expected answer doesn't match")
			}
		})
	}
}
