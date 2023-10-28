package display

import (
	"github.com/igorkichuk/tucows/common"
	"io"
)

type StringApiRes struct {
	Res string
	Err error
}

func PrintPostRes(l common.Logger, w io.Writer, res StringApiRes) {
	if res.Res != "" {
		l.Flog(w, res.Res)
	}
	if res.Err != nil {
		l.Logln(res.Err)
	}
}
