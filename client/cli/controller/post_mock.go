package controller

import "io"

type PostMock struct {
	GetImgFunc   func(grayscale bool) (io.ReadCloser, error)
	GetQuoteFunc func(key int) (string, error)
}

func (m PostMock) GetImg(grayscale bool) (io.ReadCloser, error) {
	return m.GetImgFunc(grayscale)
}

func (m PostMock) GetQuote(key int) (string, error) {
	return m.GetQuoteFunc(key)
}
