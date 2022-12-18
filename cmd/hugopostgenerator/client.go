package hugopostgenerator

import (
	"embed"
	"errors"
	"os"
	"path/filepath"
)

//go:embed templates
var fs embed.FS

type ClientOptions struct {
	PostFolder  string
	ImageFolder string
}

type Client struct {
	postFolder string
}

func NewClient(opt *ClientOptions) (*Client, error) {

	c := &Client{}

	postFolder := opt.PostFolder
	if filepath.IsAbs(postFolder) {
		return nil, errors.New("post folder path must be relative")
	}
	if _, err := os.Stat(postFolder); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("post folder must exist")
	}
	c.postFolder = postFolder

	return c, nil
}
