package mutual

import (
	"errors"

	"github.com/igorkichuk/tucows/api/quote"
)

// PostBL works with post related information
type PostBL struct {
	quoteProvider quote.Provider
}

func NewPostBL(qp quote.Provider) PostBL {
	return PostBL{
		quoteProvider: qp,
	}
}

func (bl PostBL) GetQuote(key int) (string, error) {
	if key > 999999 {
		return "", errors.New("key has to be less than 999999")
	}
	uri, err := bl.quoteProvider.Get(key)
	if err != nil {
		return "", err
	}

	return uri, nil
}
