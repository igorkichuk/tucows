package img

import (
	"fmt"
	"io"
	"net/http"
)

type picsum struct {
	client *http.Client
}

func NewPicsumImageProvider(client *http.Client) Provider {
	return picsum{
		client: client,
	}
}

type Provider interface {
	// GetImage returns io.ReadCloser. Responsibility of closing is on the caller.
	GetImage(h, w uint, gray bool) (io.ReadCloser, error)
	GetURL(h, w uint, gray bool) (string, error)
}

// GetImage returns io.ReadCloser. Responsibility of closing is on the caller.
func (p picsum) GetImage(h, w uint, gray bool) (io.ReadCloser, error) {
	data, _, err := p.get(h, w, gray)
	return data, err
}

func (p picsum) GetURL(h, w uint, gray bool) (string, error) {
	body, url, err := p.get(h, w, gray)
	if body != nil {
		_ = body.Close()
	}

	return url, err
}

func (p picsum) get(h, w uint, grayscale bool) (io.ReadCloser, string, error) {
	var imgURL = ""
	var imgURLp = &imgURL
	p.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		*imgURLp = req.URL.String()

		return nil
	}

	var grayscaleFlag string
	if grayscale {
		grayscaleFlag = "?grayscale"
	}

	uri := fmt.Sprintf("https://picsum.photos/%d/%d/%s", w, h, grayscaleFlag)

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, "", err
	}

	res, err := p.client.Do(req)
	if err != nil {
		return nil, "", err
	}
	if res.StatusCode != http.StatusOK {
		_ = res.Body.Close()
		return nil, "", fmt.Errorf("img client status code error: %d", res.StatusCode)
	}

	return res.Body, imgURL, err
}
