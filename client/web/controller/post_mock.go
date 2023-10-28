package controller

type PostMock struct {
	GetImgFunc   func(grayscale bool) (string, error)
	GetQuoteFunc func(key int) (string, error)
}

func (m PostMock) GetImg(grayscale bool) (string, error) {
	return m.GetImgFunc(grayscale)
}

func (m PostMock) GetQuote(key int) (string, error) {
	return m.GetQuoteFunc(key)
}
