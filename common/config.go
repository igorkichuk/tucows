package common

import "net/http"

type Config struct {
	ImgWidth  uint `env:"IMG_WIDTH" envDefault:"500"`
	ImgHeight uint `env:"IMG_HEIGHT" envDefault:"400"`
}

var DefaultHTTPClient = http.DefaultClient
