package splash

import (
	"bytes"
	"errors"
	"image"
	"io"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

func FetchNewest() (image.Image, error) {
	res, err := http.Get("https://xkcd.com/atom.xml")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	imgHTML, err := xpathFindString("/feed/entry[1]/summary", res.Body)
	if err != nil {
		return nil, err
	}

	imgURL, err := xpathFindString("/img/@src", bytes.NewBufferString(imgHTML))
	if err != nil {
		return nil, err
	}

	res, err = http.Get(imgURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, _, err := image.Decode(res.Body)

	return img, err
}

func xpathFindString(expr string, rd io.Reader) (string, error) {
	root, err := xmlpath.Parse(rd)
	if err != nil {
		return "", err
	}

	p := xmlpath.MustCompile(expr)
	s, ok := p.String(root)
	if !ok {
		return "", errors.New("not found")
	}
	return s, nil
}
