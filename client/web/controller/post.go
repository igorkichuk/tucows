package controller

import (
	"github.com/igorkichuk/tucows/client/mutual"
	"github.com/igorkichuk/tucows/common"
)

type qtProvider interface {
	Get(key int) (string, error)
}

type imgProvider interface {
	GetURL(h, w uint, gray bool) (string, error)
}

type PostController struct {
	mutual.PostBL
	imgProvider imgProvider
	config      common.Config
}

func NewPostController(config common.Config, qp qtProvider, ip imgProvider) PostController {
	return PostController{
		PostBL:      mutual.NewPostBL(qp),
		imgProvider: ip,
		config:      config,
	}
}

func (c PostController) GetImg(grayscale bool) (string, error) {
	url, err := c.imgProvider.GetURL(c.config.ImgHeight, c.config.ImgWidth, grayscale)
	if err != nil {
		return "", err
	}

	return url, nil
}
