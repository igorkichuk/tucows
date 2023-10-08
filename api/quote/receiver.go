package quote

import (
	"fmt"
	"io"
	"net/http"
)

type format string

const (
	TextFormat format = "text"
	HTMLFormat format = "html"
)

type QT struct {
	client     *http.Client
	respFormat format
}

func NewQuoteProvider(client *http.Client, resType format) QT {
	return QT{
		client:     client,
		respFormat: resType,
	}
}

type Provider interface {
	Get(key int) (string, error)
}

func (q QT) Get(key int) (string, error) {
	uri := fmt.Sprintf("http://api.forismatic.com/api/1.0/?method=getQuote&format=%s&lang=en&key=%d", q.respFormat, key)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}

	res, err := q.client.Do(req)
	if err != nil {
		return "", err
	}
	if res != nil {
		defer res.Body.Close()
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("quote client status code error: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
