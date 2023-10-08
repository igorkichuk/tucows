package controller

import (
	"io"

	"github.com/igorkichuk/tucows/client/mutual"
	"github.com/igorkichuk/tucows/common"
)

type qtProvider interface {
	Get(key int) (string, error)
}

type imgProvider interface {
	// GetImage returns io.ReadCloser. Responsibility of closing is on the caller.
	GetImage(h, w uint, gray bool) (io.ReadCloser, error)
}

type PostController struct {
	mutual.PostBL
	imgProvider imgProvider
	config      common.Config
}

func NewPostController(cfg common.Config, qp qtProvider, ip imgProvider) PostController {
	return PostController{
		PostBL:      mutual.NewPostBL(qp),
		imgProvider: ip,
		config:      cfg,
	}
}

// GetImg returns io.ReadCloser. Responsibility of closing is on the caller.
func (c PostController) GetImg(grayscale bool) (io.ReadCloser, error) {
	data, err := c.imgProvider.GetImage(c.config.ImgHeight, c.config.ImgWidth, grayscale)
	if err != nil {
		return nil, err
	}

	return data, nil
}
